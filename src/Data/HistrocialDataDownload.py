import time
import webbrowser
import numpy as np
import pandas as pd
from kiteconnect import KiteTicker 
from kiteconnect import KiteConnect

# login to zerodha 

KITE_API_KEY='******************* @@enter your api key here@@@**************' 
KITE_API_SECRET='**********************@@enter your api key here****************'
print(KiteConnect(KITE_API_KEY,KITE_API_SECRET).login_url())

print(" \n ")

kite = KiteConnect(KITE_API_KEY=KITE_API_KEY)
data = kite.generate_session(((str(input("ENTER full link generated in URL :- ")).split("request_token=")[1]).split("&action"))[0],KITE_API_SECRET)
kite.set_access_token(data["access_token"])
print(" \n ")
print(" \n ")

#  inputs 

while(True):
    print(" \n ")
    while(True):
        zde=str(input("Enter Symbols manually or full_list ? 'YES' for Manually and 'NO' to Pass Full_list :=")).upper()
        if zde:
            if "YES"==zde:
                break
            if "NO" in zde:
                break
    print("\n")
    if "YES" in zde:
        while(True):
            k=str(input("ENTER SYMBOL''S ex:-'sunpharma,sbin':=")).upper()
            if k:
                z=k.split(',')
                break
    if "NO" in zde:
        while(True):
            a=input( "ENTER FULL LIST OF SYMBOL's ex:-'sunpharma,sbin,adani,ongc':=").upper()
            if a:
                z=a.split(',')
                break
    while(True):
        eexchange=str(input("ENTER EXCHANGE ex:-'NSE':=")).upper()
        if eexchange:
            break
    while(True):
        time_frame=str(input("ENTER TIME_FRAME of Data ex:-'5minute \15minute \day etc' :=")).lower()
        if time_frame:
            break
    while(True):
        print(" \n ")
        zd=str(input("DO you want to change Date? 'NO'for Default and'YES'to change date:=")).upper()
        if zd:
            if "YES"==zd:
                break
            if "NO"==zd:
                break
    if "YES" in zd:
        while(True):
            sdate =str(input("Starting Date in formatof '2019-07-23' :=")).lower()
            if sdate:
                break
        while(True):
            todate =str(input("Ending Date in format of '2019-10-23' :=")).lower()
            if todate:
                break
    if "NO" in zd:
        sdate ="2019-07-23"
        todate="2019-10-23"
    print(" \n ")
    while(True):
        st=str(input("Is All above Entered inputs are TRUE ? 'YES' to continute or 'NO' to enter again ? ")).upper()
        if st:
            if "YES"==st:
                break
            else:
                pass
    if "YES"==st:
                break
tokenall=[]
symbl=[]
aa=0
print(" \n ")
print("Processing.....")

while(True):
    ttoken=int(pd.DataFrame(kite.ltp(eexchange+":"+z[aa])).iloc[-2,0])
    tokenall.append(ttoken) # fetching tokens
    symbl.append(z[aa])
    aa=aa+1
    if aa==len(z):
        break
        
print("Processing.....")        
ee=0

# downloading data

while(True):
    dff=kite.historical_data(tokenall[ee],sdate,todate,time_frame,0)
    time.sleep(1)
    dfw=pd.DataFrame(dff)
    s=f"{ symbl[ee]}.csv" # writing to csv
    dfw.to_csv(s)
    ee=ee+1
    if ee==len(z):
        print("\n")
        print("Download Complete")
        break

"""
Link this to downloaded file
"""
# Data_File = pd.DataFrame(pd.read_csv("file.csv"))
# new_file[["date", "open", "high", "low", "close"]]