Set objArgs = WScript.Arguments 
Set oFSO = CreateObject("Scripting.FileSystemObject")  
sFolder = objArgs(0)  
Set xmlDoc = _
  CreateObject("Microsoft.XMLDOM")  

For Each oFile In oFSO.GetFolder(sFolder).Files 
   if Mid(oFile.NAME,1,10) = "jasperForm" then
   oFile.Delete TRUE
end if
next

For Each oFile In oFSO.GetFolder("\\192.168.146.5\paynotes\").Files 

   if Mid(oFile.NAME,1,10) = "jasperForm" then
 
    oFile.Copy oFSO.BuildPath(sFolder, oFile.Name), True 

   
end if


Next

  
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


Set objExcel = CreateObject("Excel.Application")
Set oFSO = CreateObject("Scripting.FileSystemObject")
if oFSO.FileExists(sFolder+"sber.xml") then oFSO.DeleteFile sFolder+"sber.xml"  end if

For Each oFile In oFSO.GetFolder(sFolder).Files 
   if Mid(oFile.NAME,1,10) = "jasperForm" then
      Set objfiles = _
      xmlDoc.createElement("files")  
      objRoot.appendChild objfiles  
  
      Set objWorkbook = objExcel.Workbooks.Open(oFile)
      objExcel.Application.Visible = False
   
      Set objfilesname1 = _
      xmlDoc.createElement("number") 
      objfilesname1.text = Trim(Replace( objExcel.Cells(6, 12).Value,"¹ ",""))
      objfiles.appendChild objfilesname1 
        
      Set objfilesname2 = _
      xmlDoc.createElement("date") 
      objfilesname2.text =  objExcel.Cells(5, 15).Value
      objfiles.appendChild objfilesname2 
      
      Set objfilesname3 = _
      xmlDoc.createElement("summ") 
      objfilesname3.text =  objExcel.Cells(11, 21).Value
      objfiles.appendChild objfilesname3 
      objExcel.ActiveWorkbook.Close 

      Set objfilesname = _
      xmlDoc.createElement("filename") 
      objfilesname.text = oFSO.GetBaseName(oFile)+".xls" 
      objfiles.appendChild objfilesname    
  
  
  end If   
Next
objExcel.Application.Quit  

Set objIntro = _
  xmlDoc.createProcessingInstruction _
  ("xml","version='1.0'")  
xmlDoc.insertBefore _
  objIntro,xmlDoc.childNodes(0)    
xmlDoc.Save sFolder+"sber.xml" 
For Each oFile In oFSO.GetFolder(sFolder).Files 
   if Mid(oFile.NAME,1,10) = "jasperForm" then
   oFile.Delete TRUE
end if
next
WScript.Quit