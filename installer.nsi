; Auto-Unterhaltsrechner NSIS Installer Script
; This script creates an installer for the Auto-Unterhaltsrechner application

!define APPNAME "Auto-Unterhaltsrechner"
!define COMPANYNAME "Auto-Unterhaltsrechner"
!define DESCRIPTION "Automatischer Unterhaltsrechner für Deutschland"
!define VERSIONMAJOR 0
!define VERSIONMINOR 2
!define VERSIONBUILD 0
!define HELPURL "https://github.com/yourusername/auto-unterhaltsrechner"
!define UPDATEURL "https://github.com/yourusername/auto-unterhaltsrechner/releases"
!define ABOUTURL "https://github.com/yourusername/auto-unterhaltsrechner"

RequestExecutionLevel admin
InstallDir "$PROGRAMFILES\${APPNAME}"
LicenseData "README.md"
Name "${APPNAME}"
outFile "Auto-Unterhaltsrechner-Setup.exe"

!include LogicLib.nsh
!include "FileFunc.nsh"

page license
page directory
page instfiles

!macro VerifyUserIsAdmin
UserInfo::GetAccountType
pop $0
${If} $0 != "admin"
        messageBox mb_iconstop "Administrator rights required!"
        setErrorLevel 740
        quit
${EndIf}
!macroend

function .onInit
	setShellVarContext all
	!insertmacro VerifyUserIsAdmin
functionEnd

section "install"
	setOutPath $INSTDIR
	file "auto-unterhaltsrechner.exe"
	file "README.md"
	file "CHANGELOG.md"
	
	writeUninstaller "$INSTDIR\uninstall.exe"

	createDirectory "$SMPROGRAMS\${APPNAME}"
	createShortCut "$SMPROGRAMS\${APPNAME}\${APPNAME}.lnk" "$INSTDIR\auto-unterhaltsrechner.exe" "" ""
	createShortCut "$SMPROGRAMS\${APPNAME}\Uninstall.lnk" "$INSTDIR\uninstall.exe" "" ""
	
	createShortCut "$DESKTOP\${APPNAME}.lnk" "$INSTDIR\auto-unterhaltsrechner.exe" "" ""

	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "DisplayName" "${APPNAME} - ${DESCRIPTION}"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "UninstallString" "$\"$INSTDIR\uninstall.exe$\""
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "QuietUninstallString" "$\"$INSTDIR\uninstall.exe$\" /S"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "InstallLocation" "$\"$INSTDIR$\""
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "DisplayIcon" "$\"$INSTDIR\auto-unterhaltsrechner.exe$\""
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "Publisher" "${COMPANYNAME}"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "HelpLink" "${HELPURL}"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "URLUpdateInfo" "${UPDATEURL}"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "URLInfoAbout" "${ABOUTURL}"
	WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "DisplayVersion" "${VERSIONMAJOR}.${VERSIONMINOR}.${VERSIONBUILD}"
	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "VersionMajor" ${VERSIONMAJOR}
	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "VersionMinor" ${VERSIONMINOR}
	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "NoModify" 1
	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "NoRepair" 1

	${GetSize} "$INSTDIR" "/S=0K" $0 $1 $2
	IntFmt $0 "0x%08X" $0
	WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}" "EstimatedSize" "$0"
sectionEnd

function un.onInit
	SetShellVarContext all
	MessageBox MB_OKCANCEL "Möchten Sie ${APPNAME} wirklich deinstallieren?" IDOK next
		Abort
	next:
	!insertmacro VerifyUserIsAdmin
functionEnd

section "uninstall"
	delete "$SMPROGRAMS\${APPNAME}\${APPNAME}.lnk"
	delete "$SMPROGRAMS\${APPNAME}\Uninstall.lnk"
	rmDir "$SMPROGRAMS\${APPNAME}"
	
	delete "$DESKTOP\${APPNAME}.lnk"
	
	delete "$INSTDIR\auto-unterhaltsrechner.exe"
	delete "$INSTDIR\README.md"
	delete "$INSTDIR\CHANGELOG.md"
	delete "$INSTDIR\uninstall.exe"
	rmDir "$INSTDIR"

	DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APPNAME}"
sectionEnd