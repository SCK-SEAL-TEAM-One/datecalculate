***Settings***
Library    SelaniumLibrary
Test Setup     Open Browser
Test Teardown    Close Browser

***Variables***
${URL}    localhost:3000/​web

***Test Cases***

เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 4/01/2018 วันที่สิ้นสุด 4/06/2018
เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 27/12/1994 วันที่สิ้นสุด 4/06/2018









***Keywords***
เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 4/01/2018 วันที่สิ้นสุด 4/06/2018
    Open Browser    ${URL}    chrome
    Title Should Be    Calculate Duration Between Two Dates
    ##กรอกวันที่เริ่มต้น
    Input Text    id=startDate    4/01/2018
    ##กรอกวันที่สิ้นสุด 
    Input Text    id=startDate    4/06/2018
    ##กดปุ่มคำนวณหาผลลัพธ์  
    Click Elemnt    id=calculatedate
    ##แสดงรายละเอียดวันเริ่มต้น
    Wait Until Page Contains    Thursday, 4 January 2018
    ##แสดงรายละเอียดวันสิ้นสุด
    Wait Until Page Contains    Monday, 4 June 2018
    ##แสดงจำนวนผลลัพธ์ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    152 days
    ##แสดงระยะเวลา ปี เดือน วัน ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    5 months, 1 day including the end date
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





เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 27/12/1994 วันที่สิ้นสุด 4/06/2018
    Open Browser    ${URL}    chrome
    Title Should Be    Calculate Duration Between Two Dates
    ##กรอกวันที่เริ่มต้น
    Input Text    id=startDate    27/12/2018
    ##กรอกวันที่สิ้นสุด 
    Input Text    id=startDate    4/06/2018
    ##กดปุ่มคำนวณหาผลลัพธ์  
    Click Elemnt    id=calculatedate
    ##แสดงรายละเอียดวันเริ่มต้น
    Wait Until Page Contains    Wednesday, 4 December 1995
    ##แสดงรายละเอียดวันสิ้นสุด
    Wait Until Page Contains    Monday, 4 June 2018
    ##แสดงจำนวนผลลัพธ์ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    8196 days
    ##แสดงระยะเวลา ปี เดือน วัน ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    22 years, 5 months, 9 day including the end date
    ##แสดงระยะเวลาจำนวน วินาที ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    708,134,400 seconds
    ##แสดงระยะเวลาจำนวน นาที ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    11,802,240 minutes
    ##แสดงระยะเวลาจำนวน ชั่วโมง ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    196,704 hours
    ##แสดงระยะเวลาจำนวน วัน ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    8196 days
    ##แสดงระยะเวลาจำนวน สัปดาห์ ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    1170 weeks and 6 days
    ##แสดงระยะเวลาจำนวน ค่าเฉลี่ยในหนึ่งปี ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    2245.48% of 2018






