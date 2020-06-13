
var app= new Vue({
    el:'#app',
    data:{
        count:0,
        visible: false,
        items: [
            { message: {text:"hi"} },
            { message: {text:"bye"} }
          ],
          info: {}, 
    },
    methods:{
        handleSubmit: function Search(params) {
            params.preventDefault()
            console.log("send: ", this.movieTitle)
            axios
              .get('/s?q=' + this.movieTitle)
              .then(response => {
                this.info = response.data;
                console.log("data", this.info.count);
                console.log("inf", this.info.filminfo.length);
                mass = "" 
                cd='<div>'
                for (let i = 0; i < this.info.filminfo.length; i++) { // выведет 0, затем 1, затем 2
                  console.log("len:", i, ") ", this.info.filminfo.length)
                  mass = mass +genItem(this.info.filminfo[i])
      
                }
                this.message = "Result for: " + this.movieTitle +  '<div class="result">'+'<div class="resultItem"> ' + mass + '</div>'+ '</div>'
              })
              .catch(error => {
                console.log(error);
                this.errored = true;
              })
              .finally(() => (this.loading = false));
              this.visible=  true
          },
        click: function(){
            this.count= this.count+1
            this.items.push({message: {text:"hi"}} , {message: {text:"add"}});

            const [first,t,e] =genItem()
            console.log("data  returned", first,t,e);
        }
    },
   




});
