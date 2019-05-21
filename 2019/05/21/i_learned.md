## 엑셀 셀 병합 매크로 

출처 : https://fillin.tistory.com/76
-==================
Sub MergeMacro()

' 선택 영역에서 인접 셀에 같은 값이 있는 경우 셀을 병합함'
If Selection.Cells.Count < 2 Then
    MsgBox "작업할 범위를 먼저 선택하세요"

Exit Sub

End If

Dim iRow As Integer, iCol As Integer, tR As Integer, tC As Integer, sVal As String
Dim rMax As Integer, cMax As Integer, iCount As Integer, cSave As Integer

Application.DisplayAlerts = False
Application.ScreenUpdating = False

iRow = Selection.Cells(1).Row: iCol = Selection.Cells(1).Column

cSave = iCol

rMax = Selection.Cells(Selection.Cells.Count).Row
cMax = Selection.Cells(Selection.Cells.Count).Column

tR = 0: tC = 0: iCount = 0

Do While iRow <= rMax
    sVal = Cells(iRow, iCol)
    ' 현재 셀이 병합 셀이 아닌경우'
    If Cells(iRow, iCol).Cells.Count = 1 And Trim(Cells(iRow, iCol)) <> "" Then
    ' 우측 연속 셀 검사'
    Do While Cells(iRow, iCol + tC + 1) = sVal
        tC = tC + 1
    Loop
    
    If tC > 0 Then
    ' 우측 병합대상 있는 경우'
        Do While Cells(iRow + tR + 1, iCol) = sVal
    For i = 0 To tC
        If Cells(iRow + tR + 1, iCol + i) <> sVal Then Exit Do
        
    Next i
    
    tR = tR + 1
    
    Loop
    Range(Cells(iRow, iCol), Cells(iRow + tR, iCol + tC)).Merge
    iCol = iCol + tC
    iCount = iCount + 1
Else
    Do While Cells(iRow + tR + 1, iCol) = sVal
        tR = tR + 1
    Loop
    If tR > 0 Then
        Range(Cells(iRow, iCol), Cells(iRow + tR, iCol)).Merge
        iCount = iCount + 1
    End If
    iCol = iCol + 1
End If
tC = 0: tR = 0
Else
    Cells(iRow, iCol).Offset(0, 1).Select
    iCol = Selection.Column
End If

If iCol > cMax Then
    iCol = cSave: iRow = iRow + 1
End If

Loop

Application.DisplayAlerts = True
Application.ScreenUpdating = True

MsgBox Trim(iCount) & "개의 병합셀이 만들어졌습니다."

End Sub

