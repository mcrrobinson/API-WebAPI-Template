function StartEventStream(params) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4) {
        if(this.status == 200){
          console.log("Stream has been started.");
        } else if(this.status == 201) {
          document.getElementById("modal-body-message").innerHTML = "The stream was already on.";
          $('#toastModal').modal('show');
          setTimeout( function(){ 
            $('#toastModal').modal('hide');
          },2000);
        }
      }
    };
    xhttp.open("GET", "http://localhost:8080/startStream", true);
    xhttp.send();
}