var xtable = document.getElementsByClassName("x-table")[0];
var tbo = xtable.children[1];

for(var i = 0; i < tbo.children.length; i++){
    tbo.children[i].children[3].innerText = ""
}