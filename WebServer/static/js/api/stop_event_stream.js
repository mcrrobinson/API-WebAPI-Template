function StopEventStream(params) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4) {
        if(this.status == 200){
          console.log("Stream has been closed.");
        } else if(this.status == 202) {
          document.getElementById("modal-body-message").innerHTML = "The stream was already closed.";
          $('#toastModal').modal('show');
          setTimeout( function(){ 
            $('#toastModal').modal('hide');
          },2000);
        }
      }
    };
    xhttp.open("GET", "http://localhost:8080/stopStream", true);
    xhttp.send(); 
}