Set objArgs = WScript.Arguments

Set xmlDoc = _
  CreateObject("Microsoft.XMLDOM")  
  
Set objRoot =  xmlDoc.createElement("xslfiles")    

Set xmlns = xmlDoc.createAttribute("version")
    xmlns.text = "1"
    objRoot.setAttributeNode xmlns
    
xmlDoc.appendChild objRoot     

Set objRecord = _
    xmlDoc.createElement("code") 
    objRecord.Text = "0"
objRoot.appendChild objRecord 

Set objRecord = _
    xmlDoc.createElement("message") 
    objRecord.Text = ""
objRoot.appendChild objRecord

sFolder = objArgs(0)  
Set objExcel = CreateObject("Excel.Application")
Set oFSO = CreateObject("Scripting.FileSystemObject")

if oFSO.FileExists(sFolder+"zapsib.xml") then oFSO.DeleteFile sFolder+"zapsib.xml"  end if  




For Each oFile In oFSO.GetFolder(sFolder).Files
    if Mid(oFile.NAME,1,5) = "Sheet" then
    Set objfiles = _
    xmlDoc.createElement("files")  
    objRoot.appendChild objfiles
      
    Set objWorkbook = objExcel.Workbooks.Open(oFile)
    objExcel.Application.Visible = False
    
    If  objExcel.Cells(3, 1).Value="œÀ¿“≈∆ÕŒ≈ œŒ–”◊≈Õ»≈"  then   
       Set objfilesname1 = _
       xmlDoc.createElement("number") 
       objfilesname1.text = Trim(Replace( objExcel.Cells(3, 4).Value,"π ",""))
       objfiles.appendChild objfilesname1                      
  
       Set objfilesname2 = _
       xmlDoc.createElement("date") 
       objfilesname2.text =  objExcel.Cells(3, 6).Value
       objfiles.appendChild objfilesname2  

       Set objfilesname3 = _
       xmlDoc.createElement("summ") 
       objfilesname3.text =  objExcel.Cells(8, 7).Value
       objfiles.appendChild objfilesname3 
       objExcel.ActiveWorkbook.Close 

       Set objfilesname = _
       xmlDoc.createElement("filename") 
       objfilesname.text = "\"+objArgs(1)+"\"+oFSO.GetBaseName(oFile)+".xls"
       objfiles.appendChild objfilesname  
  
  
  end If  
end if
Next



objExcel.Application.Quit


Set objIntro = _
  xmlDoc.createProcessingInstruction _
  ("xml","version='1.0'")  
  xmlDoc.insertBefore _
  objIntro,xmlDoc.childNodes(0)  

xmlDoc.Save sFolder+"zapsib.xml" 

Set oFSO = CreateObject("Scripting.FileSystemObject")

if not oFSO.FolderExists("\\192.168.146.5\paynotes\"+objArgs(1)+"\") then oFSO.CreateFolder("\\192.168.146.5\paynotes\"+objArgs(1)+"\")



For Each oFile In oFSO.GetFolder(sFolder).Files 

   
    
    oFile.Copy oFSO.BuildPath("\\192.168.146.5\paynotes\"+objArgs(1), oFile.Name), True 

next

WScript.Quit