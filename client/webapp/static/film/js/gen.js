function genItem(object){
  country=""
  creators=""
  for (let i = 0; i < object.directors.length; i++) { // выведет 0, затем 1, затем 2
    console.log("creators:", i, ") ",object.directors[i].name)
    creators = creators +" "+object.directors[i].name

  }
  for (let i = 0; i < object.countries.length; i++) { // выведет 0, затем 1, затем 2
    console.log("country:", i, ") ",object.countries[i].name)
    country = country +" "+object.countries[i].name

  }
  stor=""
  if (object.short_story.length>400) {
    stor=object.short_story.substr(0, 400)+'<details><summary><span><b>Раскрыть</b></span></summary><span>'+object.short_story.substr(400, object.short_story.length)+'</span></details>'
} else{
stor=object.short_story 
}
 console.log(stor)
 
pic='<a href="/film/'+object.id+'"><img src="'+object.poster+'" width="110" height="165" alt="logo.png"></a>' 
info='<div><b>Название: </b>'+'<a href="/film/'+object.id+'">'+object.title+'</a>'+'</div>'+'<div><b>Оригинальное название: </b>'+object.original_title+'</div><div><b>Год: </b>'+object.year+'</div><div><b>Страна: </b>'+country+'</div>'+'<div><b>Создатели: </b>'+creators+'</div>'+'<div><b>Описание: </b>'+stor+'</div>'
    result='<div class="filmItem" >'+pic+'</div>'+'<div class="filmItem" >'+info+'</div>'



   res=' <div class="item" ><div class="filmbox" >'+ result+'</div></div>'
    console.log("gen")
   
    return res
  }