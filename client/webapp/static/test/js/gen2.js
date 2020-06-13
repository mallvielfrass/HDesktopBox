function genItem (){
    res="hi"
    var xhr = new XMLHttpRequest();

    // 2. Конфигурируем его: GET-запрос на URL 'phones.json'
    xhr.open('GET', '/j?q=', false);
    
    // 3. Отсылаем запрос
    xhr.send();
    
    // 4. Если код ответа сервера не 200, то это ошибка
    if (xhr.status != 200) {
      // обработать ошибку
      alert( xhr.status + ': ' + xhr.statusText ); // пример вывода: 404: Not Found
    } else {
      // вывести результат
     // alert( xhr.responseText ); // responseText -- текст ответа.
    }  
 // console.log(xhr.responseText)
  var jsonResponse = JSON.parse(xhr.responseText);
 // console.log( jsonResponse)
// inf=  {message: {text:"hi"}} , {message: {text:"bye"}} 
          return  jsonResponse ;
        };
        
        
 