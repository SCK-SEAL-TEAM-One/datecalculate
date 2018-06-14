***Settings***
Library    SeleniumLibrary
Test Teardown    Close All Browsers
           

***Variables***
${URL}    https://localhost:3000/​web
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
    ##แสดงรายละเอียดวันเริ่มต้น
    Wait Until Page Contains    Thursday, 4 January 2018
    ##แสดงรายละเอียดวันสิ้นสุด
    Wait Until Page Contains    Monday, 4 June 2018
    ##แสดงจำนวนผลลัพธ์ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    152 days
    ##แสดงระยะเวลา ปี เดือน วัน ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    5 months, 1 day
    ##แสดงระยะเวลาจำนวน วินาที ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    13,132,800 seconds
    ##แสดงระยะเวลาจำนวน นาที ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    218,880 minutes
    ##แสดงระยะเวลาจำนวน ชั่วโมง ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    3,648 hours
    ##แสดงระยะเวลาจำนวน วัน ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    152 days
    ##แสดงระยะเวลาจำนวน สัปดาห์ ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    21 weeks and 5 days
    ##แสดงระยะเวลาจำนวน ค่าเฉลี่ยในหนึ่งปี ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    41.64% of 2018



