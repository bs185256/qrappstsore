
      var form = document.getElementById('myform');
      var usnname = document.getElementById("usnname");
      var pwd = document.getElementById("pwd");
    let url = "https://gateway-staging.ncrcloud.com/security/authentication/login";
    //   const username = usnname.value.trim();
    //     const password = pwd.value.trim();

  function authenticate(){
       getToken();
    console
        alert("Login failed! Please try again.");
            return false;
        
      }

let encoded = window.btoa('NexusApp.1');
let auth = 'Basic ncrqrapp:' + encoded;
var request = new Request(url, {
    method: 'POST',
    cache: 'no-cache',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': auth,
      'nep-application-key': '8a808f0d65ba2ef401671f2e728402fa',
      'nep-organization': '/'
    },
    mode: 'no-cors',
    redirect: 'follow',
    referrerPolicy: 'no-referrer'
  });
  
  async function getToken() {
    const response = await fetch(request)
    .then(response => response.json())
    .then(
      console.log('Success:')
    )
    .catch((error) => {
      console.log('Error:', error);
    });
    const data = await response.json();
    return true;
  }

  
  