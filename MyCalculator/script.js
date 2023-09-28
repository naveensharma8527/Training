function clearScreen() {
    document.getElementById("display").value = "";
}


function display(val) {
    document.getElementById("display").value += val;
    console.log(val);
}
 

function cal() {
    var p = document.getElementById("display").value;
    var q = eval(p);
    document.getElementById("display").value = q;
}