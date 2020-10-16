if (typeof (EventSource) !== "undefined") {
    var source = new EventSource("http://localhost:3000");
    source.onmessage = function (event) {
      objectThing = JSON.parse(event.data);
      $('#example').dataTable().fnAddData([
        objectThing["name"],
        objectThing["position"],
        objectThing["office"],
        objectThing["age"],
        objectThing["startdate"],
        "£"+objectThing["salary"],
      ]);
      valueChartCounter++;
    };
} else {
    document.getElementById("modal-body-message").innerHTML = "Your browser does not support EventSources, this means it can't receieve any information from the API.";
    $('#toastModal').modal('toggle');
}