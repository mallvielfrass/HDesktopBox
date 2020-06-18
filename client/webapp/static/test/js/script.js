
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    movieTitle: "",
    info: {},

    render: false,
  },
  methods: {
    handleSubmit: function Search(params) {
      axios.get('/j?q=' + this.movieTitle).then(response => {
        this.info = genItem2(response.data)
        this.render = true
        console.log("send: ", this.movieTitle)
       // console.log("rec: ", this.info)
        console.log("rec: ", this.info[0].film.title)
      })
     
    

    },

    handleChange(event) {
      this.movieTitle = event.target.value
     // console.log(this.movieTitle)
    }
  },
})
