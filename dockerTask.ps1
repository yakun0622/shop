<#
.SYNOPSIS
Builds and runs a Docker image.
.PARAMETER Compose
Runs docker-compose.
.PARAMETER Build
Builds a Docker image.
.PARAMETER Clean
Removes the image cb and kills all containers based on that image.
.PARAMETER Environment
The enviorment to build for (Debug or Release), defaults to Debug
.EXAMPLE
C:\PS> .\dockerTask.ps1 -Build
Build a Docker image named cb
#>

Param(
    [Parameter(Mandatory=$True,ParameterSetName="Compose")]
    [switch]$Compose,
    [Parameter(Mandatory=$True,ParameterSetName="Build")]
    [switch]$Build,
    [Parameter(Mandatory=$True,ParameterSetName="Clean")]
    [switch]$Clean,
    [parameter(ParameterSetName="Compose")]
    [parameter(ParameterSetName="Build")]
    [parameter(ParameterSetName="Clean")]
    [ValidateNotNullOrEmpty()]
    [String]$Environment = "Debug"
)

$imageName="cb"
$projectName="cb"
$publicPort=8080
$url="http://localhost:$publicPort"

# Kills all running containers of an image and then removes them.
function CleanAll () {
    $composeFileName = "docker-compose.yml"
    if ($Environment -ne "Release") {
        $composeFileName = "docker-compose.$Environment.yml"
    }

    if (Test-Path $composeFileName) {
        docker-compose -f "$composeFileName" -p $projectName down --rmi all

        $danglingImages = $(docker images -q --filter 'dangling=true')
        if (-not [String]::IsNullOrWhiteSpace($danglingImages)) {
            docker rmi -f $danglingImages
        }
    }
    else {
        Write-Error -Message "$Environment is not a valid parameter. File '$composeFileName' does not exist." -Category InvalidArgument
    }
}

# Builds the Docker image.
function BuildImage () {
    $composeFileName = "docker-compose.yml"
    if ($Environment -ne "Release") {
        $composeFileName = "docker-compose.$Environment.yml"
    }

    if (Test-Path $composeFileName) {
        Write-Host "Building the image $imageName ($Environment)."
        docker-compose -f $composeFileName -p $projectName build
    }
    else {
        Write-Error -Message "$Environment is not a valid parameter. File '$composeFileName' does not exist." -Category InvalidArgument
    }
}

# Runs docker-compose.
function Compose () {
    $composeFileName = "docker-compose.yml"
    if ($Environment -ne "Release") {
        $composeFileName = "docker-compose.$Environment.yml"
    }

    if (Test-Path $composeFileName) {
        Write-Host "Running compose file $composeFileName"
        docker-compose -f $composeFileName -p $projectName kill
        docker-compose -f $composeFileName -p $projectName up -d
    }
    else {
        Write-Error -Message "$Environment is not a valid parameter. File '$dockerFileName' does not exist." -Category InvalidArgument
    }
}

# Opens the remote site
function OpenSite () {
    Write-Host "Opening site" -NoNewline
    $status = 0

    #Check if the site is available
    while($status -ne 200) {
        try {
            $response = Invoke-WebRequest -Uri $url -Headers @{"Cache-Control"="no-cache";"Pragma"="no-cache"} -UseBasicParsing
            $status = [int]$response.StatusCode
        }
        catch [System.Net.WebException] { }
        if($status -ne 200) {
            Write-Host "." -NoNewline
            Start-Sleep 1
        }
    }

    Write-Host
    # Open the site.
    Start-Process $url
}

$Environment = $Environment.ToLowerInvariant()

# Call the correct function for the parameter that was used
if($Compose) {
    Compose
    OpenSite
}
elseif($Build) {
    BuildImage
}
elseif ($Clean) {
    CleanAll
}