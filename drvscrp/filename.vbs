Set objArgs = WScript.Arguments
Set oFSO = CreateObject("Scripting.FileSystemObject")

dstFolder=objArgs(1)

newFolder=objArgs(2)
wscript.echo "start copy"  

if (oFSO.FileExists(dstFolder+newFolder+"\RurPaymentDemand.xls")) then 
oFSO.DeleteFile(dstFolder+newFolder+"\RurPaymentDemand.xls"),true  
end if  




For Each oFile In oFSO.GetFolder("\\192.168.146.5\paynotes\").Files
   if oFile.NAME = "RurPaymentDemand.xls" then        
     oFile.Copy oFSO.BuildPath(dstFolder+newFolder+"\", oFile.Name), True   
  end if
next
wscript.echo "prepare plugin"
Set objExcel = CreateObject("Excel.Application")
objExcel.Application.Visible = False
objExcel.Application.AlertBeforeOverwriting = False
objExcel.Application.DisplayAlerts = False  
Set objWorkbook = objExcel.Workbooks.Open(objArgs(0)+"1.xls")
Set objWorkbook = objExcel.Workbooks.Open(dstFolder+newFolder+"\RurPaymentDemand.xls")
wscript.echo "start plugin"  
if (oFSO.FileExists(dstFolder+"RurPaymentDemand.xls")) then 
oFSO.DeleteFile(dstFolder+"RurPaymentDemand.xls"),true  
end if  
objExcel.Application.Run "1.xls!Copy_Sheet_File",objArgs(2)   

objExcel.Application.Quit   

wscript.echo "finish"  
WScript.Quit