
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    movieTitle: "",
    info: {},
  },
  methods: {
    handleSubmit: function Search(params) {
      params.preventDefault()
      console.log("send: ", this.movieTitle)
      axios
       // .get('/s?q=' + this.movieTitle)
       .get('/j?q=' + this.movieTitle)
        .then(response => {
          this.info = response.data;
          console.log("data", this.info.count);
          console.log("info", this.info);
          console.log("inf", this.info.filminfo.length);
        //  mass = "" 
         // cd='<div>'
        //  for (let i = 0; i < this.info.filminfo.length; i++) { // выведет 0, затем 1, затем 2
         //   console.log("len:", i, ") ", this.info.filminfo.length)
         //   mass = mass +genItem(this.info.filminfo[i])

        //  }
        //  this.message = 
        })
          .catch(error => {
          console.log(error);
          this.errored = true;
        })
        .finally(() => (this.loading = false));
      
    },

    handleChange(event) {
      this.movieTitle = event.target.value
      console.log(this.movieTitle)
    }
  },
})
