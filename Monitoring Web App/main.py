import MySQLdb
import os

from flask import Flask
from flask_prometheus import monitor, mysql_fetchall


#flask : start the http server 
app = Flask(__name__)

#database
# initialize the connection to the data base 
db=MySQLdb.connect(host=os.getenv("MYSQL_HOST", "localhost"),
                   user=os.getenv("MYSQL_USER", "root"),
                   passwd=os.getenv("MYSQL_PASSWORD", ""),
                   db=os.getenv("MYSQL_DB", "app"))
                 
# routing 


@app.route('/')
def index():
  return "Flask is up and running \n"
@app.route('/query')
def query():
  res = mysql_fetchall(db, "select 1 as id, RAND()*10 as rand union select 2 as id, RAND()*100 as rand")
  ouput = ""
  for row in res:
    line = ",".join(str(s) for s in row)
    output += line+"\n"
  return output
  
@app.route('/sleep')
def sleep():
  res = mysql_fetchall(db, "select SLEEP(RAND()*10) as sleeping")
  return "Done!"
  
#monitoring 

monitor(app, port=8000)

#run app 
app.run(host="172.16.2.30")
