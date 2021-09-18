function show(){
  var x = document.getElementById("hi3");
  y = 0
  for (i = 0; i <= x.options.length; i++){
    x.options[i] = null;
  }
  
  for (i = 0 ; i < document.getElementById(document.getElementById("choice").value+"1").childElementCount; i++){
    var option = document.createElement("option");
    option.text = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
    option.value = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
    console.log(option.text);
    console.log(i);
    x.add(option);
    y = i + 1;
  }

  var z = document.getElementById("hi3");
  while(z.options.length - y != 0){
    z.options[0] = null
    console.log("the length is", z.options.length);
  }
  text();
}

function text(){
  if(document.getElementById(document.getElementById("docschange").value) == null){
    document.getElementById("fname").innerHTML = "";
  }else {
    document.getElementById("fname").innerHTML = document.getElementById(document.getElementById("docschange").value).innerHTML;
  }

  if(document.getElementById("docschange") == null){
    document.getElementById("changename1").innerHTML = "";
    document.getElementById("fname").innerHTML = "";
  }else {
    document.getElementById("changename1").value = document.getElementById("docschange").value;
  }

  if(document.getElementById(document.getElementById("docschange").value + "1") == null){
    document.getElementById("changename2").innerHTML = "";
  }else {
    let str = document.getElementById("changename2").value = document.getElementById(document.getElementById("docschange").value + "1").src;
    str = str.split("/preview")[0]
    str = str.replace(".png","")
    document.getElementById("changename2").value = str.split("/images/")[1]
  }
  myCanvas()
}

function change1(){
  
  console.log(document.getElementById("action").value);
  if(document.getElementById("action").value == "add"){
    document.getElementById("append").style.display = "block";
    document.getElementById("thingadd").innerHTML = "add"
  }else{
    document.getElementById("append").style.display = "change";
  }

  if(document.getElementById("action").value == "remove"){
    document.getElementById("remove").style.display = "block";
    document.getElementById("thingadd").innerHTML = "change"
  }else{
    document.getElementById("remove").style.display = "none";
  }

  if(document.getElementById("action").value == "change"){
    document.getElementById("change").style.display = "block";
    document.getElementById("thingadd").innerHTML = "change";
  }else{
    document.getElementById("change").style.display = "none";
  }
}

/*function show1(){
  if(document.getElementById("action").value == "change"){
    
    document.getElementById("change").style.display = "block";

    var x = document.getElementById("docschange");
    y = 0
    for (i = 0; i <= x.options.length; i++){
      x.options[i] = null;
    }
    
    for (i = 0 ; i < document.getElementById(document.getElementById("choice").value+"1").childElementCount; i++){
      var option = document.createElement("option");
      option.text = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
      option.value = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
      console.log(option.text);
      console.log(i);
      x.add(option);
      y = i + 1;
    }
  
    var z = document.getElementById("docschange");
    while(z.options.length - y != 0){
      z.options[0] = null
      console.log("the length is", z.options.length);
    }
    text();
  }else{
    document.getElementById("change"). style.display = "none";
  }
  
  if(document.getElementById("action").value == "remove"){
   
    document.getElementById("remove"). style.display = "block";

    var x = document.getElementById("docsremove");
    y = 0
    
    for (i = 0; i <= x.options.length; i++){
      x.options[i] = null;
    }
    
    for (i = 0 ; i < document.getElementById(document.getElementById("choice").value+"1").childElementCount; i++){
      var option = document.createElement("option");
      option.text = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
      option.value = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
      console.log(option.text);
      console.log(i);
      x.add(option);
      y = i + 1;
    }
  
    var z = document.getElementById("docsremove");
    
    while(z.options.length - y != 0){
      z.options[0] = null
      console.log("the length is", z.options.length);
    }
 
  }else{
    document.getElementById("remove"). style.display = "none";
  }
}*/

function showchange(){
  var x = document.getElementById("docschange");
  y = 0
  for (i = 0; i <= x.options.length; i++){
    x.options[i] = null;
  }
  
  for (i = 0 ; i < document.getElementById(document.getElementById("choicechange").value+"1").childElementCount; i++){
    var option = document.createElement("option");
    option.text = document.getElementById(document.getElementById("choicechange").value+"1").children[i].innerHTML;
    option.value = document.getElementById(document.getElementById("choicechange").value+"1").children[i].innerHTML;
    console.log(option.text);
    console.log(i);
    x.add(option);
    y = i + 1;
  }

  var z = document.getElementById("docschange");
  while(z.options.length - y != 0){
    z.options[0] = null
    console.log("the length is", z.options.length);
  }
  text();

}

function showremove(){
  var x = document.getElementById("docsremove");
  y = 0
  
  for (i = 0; i <= x.options.length; i++){
    x.options[i] = null;
  }
  
  for (i = 0 ; i < document.getElementById(document.getElementById("choice").value+"1").childElementCount; i++){
    var option = document.createElement("option");
    option.text = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
    option.value = document.getElementById(document.getElementById("choice").value+"1").children[i].innerHTML;
    console.log(option.text);
    console.log(i);
    x.add(option);
    y = i + 1;
  }

  var z = document.getElementById("docsremove");
  
  while(z.options.length - y != 0){
    z.options[0] = null
    console.log("the length is", z.options.length);
  }
}

function myCanvas() {
  var c = document.getElementById("myCanvas");
  var ctx = c.getContext("2d");
  //ctx.clearRect(0, 0, c.width, c.height)
  var img = document.getElementById(document.getElementById("docschange").value + "1");
  ctx.drawImage(img,10,10);
}