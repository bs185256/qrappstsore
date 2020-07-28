
      var form = document.getElementById('myform');
      var usnname = document.getElementById('usnname');
      var pwd = document.getElementById('pwd');
      var abc = document.getElementById()
      var objPeople = [
        { 
            username: "david",
            password: "nexus"
        },
        { 
            username: "sai",
            password: "nexus"
        },
        { 
            username: "bhavya",
            password: "nexus"
        }
    
    ];

      function authenticate(){
        debugger;
        var username = usnname.value.trim();
        var password = pwd.value.trim();
        for(var i = 0; i < objPeople.length; i++) {
            if(username == objPeople[i].username && password == objPeople[i].password) {
                console.log(username + " is logged in!!!") 
                return true;
            }
        }
        alert("Login failed! Please try again.");
            return false;
        
      }

  