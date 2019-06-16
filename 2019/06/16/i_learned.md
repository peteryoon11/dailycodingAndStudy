from flask import Flask ,Response,request
#from flask_restful import Resource, Api
#import ValidationModule
#import DBConnectionModule
import json
app = Flask(__name__)
#app.run(host='0.0.0.0', port=8787, debug=True) 
@app.route("/root",methods=["PUT"])
def root():
    print(request.json['ID'])
    print("test")
    return "test"
    


https://stackoverflow.com/questions/41472049/502-bad-gateway-with-nginx-google-app-engine-node-js
