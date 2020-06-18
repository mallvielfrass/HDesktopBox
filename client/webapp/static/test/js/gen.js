
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
