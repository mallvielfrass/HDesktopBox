
function genItem2(object){
  info = []
  jsv = object  
  console.log('len: ',  Object.keys(jsv.filminfo).length)
  for (var i = 0; i < Object.keys(jsv.filminfo).length; i++) {
    console.log('iter i:', i)
  
    country = ''
    creators = ''
   console.log("film:",jsv.filminfo[i].film.title)
    console.log("countries:", Object.keys(jsv.filminfo[i].film.countries).length)
    for (let z = 0; z <  Object.keys(jsv.filminfo[i].film.countries).length; z++) {

      console.log('country:', z, ') ', jsv.filminfo[i].film.countries[z].name)
      country = country + ' ' + jsv.filminfo[i].film.countries[z].name
    }
  
  
    for (let n = 0; n <  Object.keys(jsv.filminfo[i].film.directors).length; n++) {
 
      console.log('directors:', n, ') ', jsv.filminfo[i].film.directors[n].name)
      creators = creators + ' ' + jsv.filminfo[i].film.directors[n].name
    }
    info.push({ film: {
       idKinopoisk: jsv.filminfo[i].film.idKinopoisk,
       slogan: jsv.filminfo[i].film.slogan,
       title: jsv.filminfo[i].film.title,
       countries:country,
       poster: jsv.filminfo[i].film.poster,
       directors:creators,
       short_story: jsv.filminfo[i].film.short_story,
       year: jsv.filminfo[i].film.year,
       original_title: jsv.filminfo[i].film. original_title,
       } });
  }
  console.log(info)
  console.log(' info.length:', Object.keys(info).length)
    return  info
}


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