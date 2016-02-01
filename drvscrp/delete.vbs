Set objArgs = WScript.Arguments
Set oFSO = CreateObject("Scripting.FileSystemObject")
sDirectoryPath =objArgs(0)
sDirectoryPath ="\"+sDirectoryPath
 
wscript.echo "start delete" 
           
                set oFolder = oFSO.GetFolder(sDirectoryPath)
                set oFolderCollection = oFolder.SubFolders
                set oFileCollection = oFolder.Files
 
                For each oFile in oFileCollection
                               oFile.Delete(True)
                Next
 
                For each oDelFolder in oFolderCollection
                                oDelFolder.Delete(True)
                Next
 
                Set oFSO = Nothing
                Set oFolder = Nothing
                Set oFileCollection = Nothing
                Set oFile = Nothing         

wscript.echo "finish"  
WScript.Quit