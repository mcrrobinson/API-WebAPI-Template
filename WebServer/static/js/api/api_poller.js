// Ajax Poller.
var seconds = 0;
var el = document.getElementById('pollingCounter');
var taskQueue=document.getElementById("taskQueue");
var node;
var textNode;
(function poll() {
  $.ajax({
      url: "http://localhost:8080/startPolling",
      type: "GET",
      success: function(data) {
        taskQueue.innerHTML = "";
        data.forEach(x => {
          node = document.createElement("p");
          textNode = document.createTextNode("Task: "+x.task+" | Status: "+x.status);
          node.appendChild(textNode);
          taskQueue.appendChild(node);
          seconds = 0;
        });
      },
      error: function(data) {
        console.error(data);
      },
      dataType: "json",
      complete: setTimeout(function() {poll()}, 1000)
  })
})();

// Timer at the bottom, to show when the poller was last updated.
function incrementSeconds() {
  seconds += 0.5;
  el.innerText = "Updated " +seconds+ " seconds ago.";
}
var cancel = setInterval(incrementSeconds, 500);