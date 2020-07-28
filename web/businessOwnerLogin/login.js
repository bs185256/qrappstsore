
      const pass_field = document.querySelector('.pass-key');
      const showBtn = document.querySelector('.show');
      var form = document.getElementById('myform');
      var usnname = document.getElementById('usnname');
      var pwd = document.getElementById('password');
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
        var username = usnname.value;
        var password = pwd.value;
        for(var i = 0; i < objPeople.length; i++) {
            // check is user input matches username and password of a current index of the objPeople array
            if(username == objPeople[i].username && password == objPeople[i].password) {
                console.log(username + " is logged in!!!") 
                return true;
            }
        }
        alert("Login failed! Please try again.");
            return false;
        
      }

  