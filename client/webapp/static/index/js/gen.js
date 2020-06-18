
function genItem2(object){
  info = []
  jsv = object  
  console.log('len: ',  Object.keys(jsv.filminfo.film).length)
  for (var i = 0; i < Object.keys(jsv.filminfo.film).length; i++) {
    console.log('iter i:', i)
    console.log("film:", jsv.filminfo.film[0])
    console.log("countries:", Object.keys(jsv.filminfo.film[i].countries).length)
    country =""
    creators = ""
  
    for (let z = 0; z <  Object.keys(jsv.filminfo.film[i].countries).length; z++) {

      console.log('country:', z, ') ', jsv.filminfo.film[i].countries[z].name)
      country = country + ' ' + jsv.filminfo.film[i].countries[z].name
    }
  
  
    for (let n = 0; n <  Object.keys(jsv.filminfo.film[i].directors).length; n++) {
 
      console.log('directors:', n, ') ', jsv.filminfo.film[i].directors[n].name)
      creators = creators + ' ' + jsv.filminfo.film[i].directors[n].name
    }
    info.push({ film: {
       id: jsv.filminfo.film[i].id,
       slogan: jsv.filminfo.film[i].slogan,
       title: jsv.filminfo.film[i].title,
       countries:country,
       poster: jsv.filminfo.film[i].poster,
       directors:creators,
       short_story: jsv.filminfo.film[i].short_story,
       year: jsv.filminfo.film[i].year,
       original_title: jsv.filminfo.film[i]. original_title,
       } });
  }
 // console.log(info)
  console.log(' info.length:', Object.keys(info).length)
    return  info
}
