***Settings***
Library    SeleniumLibrary
Test Teardown    Close All Browsers
           

***Variables***
${URL}    http://localhost:3000/web
${Browser}    chrome

***Test Cases***
เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 4/01/2018 วันที่สิ้นสุด 4/06/2018
    Open Browser    ${URL}    ${Browser}
    Title Should Be    Calculate Duration Between Two Dates
    ##กรอกวันที่เริ่มต้น
    Input Text    id=startDay    4
    ##กรอกเดือนเริ่มต้น     
    Input Text    id=startMonth   1
    ##กรอกปีเริ่มต้น     
    Input Text    id=startYear    2018 
    ##กรอกวันที่สิ้นสุด
    Input Text    id=endDay    4
    ##กรอกเดือนสิ้นสุด     
    Input Text    id=endMonth   6
    ##กรอกปีสิ้นสุด     
    Input Text    id=endYear    2018 
    ##กดปุ่มคำนวณหาผลลัพธ์  
    Click Element    id=calculate
    ##แสดงจำนวนผลลัพธ์ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    152 days
