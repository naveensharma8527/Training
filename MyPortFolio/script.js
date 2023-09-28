 // Get a reference to the <h1> element
 var header = document.getElementById("heading");

 // Add a click event listener to the <h1> element
let  flag = true;
 header.addEventListener("click", function() {
     // Change the color to green
     console.log("Hello");

     if(flag){
        header.style.color = "green";
        flag = false;
     }else{
        header.style.color = "red";
        flag = true;
     }
 });