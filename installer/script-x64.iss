#define AppName "Lux"
#define Version "1.1.0"
#define Publisher "BanDev"
#define URL "https://bandev.uk/lux"
#define ExeName "lux.exe"

[Setup]
AppId={{E17378AA-71AA-4EFC-8B09-B9A18F395B2F}
AppName={#Name}
AppVersion={#Version}
AppPublisher={#Publisher}
SetupIconFile=src/lux.ico
UninstallIconFile=src/lux.ico
UninstallDisplayName={#AppName} (x64)
AppPublisherURL={#URL}
DefaultDirName={autopf}\{#AppName}
DisableDirPage=yes
DefaultGroupName={#AppName}
DisableProgramGroupPage=yes
LicenseFile=src/LICENSE.txt
PrivilegesRequired=lowest
OutputDir=/
OutputBaseFilename=setup-x64
Compression=lzma
SolidCompression=yes
WizardStyle=modern
ChangesEnvironment=yes
ArchitecturesAllowed=x64
ArchitecturesInstallIn64BitMode=x64

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "src\x64\{#ExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "src\LICENSE.txt"; DestDir: "{app}"; Flags: ignoreversion

[Registry]
Root: HKCU; Subkey: "Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; \
    Check: NeedsAddPath(ExpandConstant('{app}'))

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKEY_CURRENT_USER,
    '\Environment',
    'Path', OrigPath)
  then begin
    Result := True;
    exit;
  end;
  Result := Pos(';' + Param + ';', ';' + OrigPath + ';') = 0;
end;

