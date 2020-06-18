
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    movieTitle: "",
    info: {},
    status: false,
    render: false,
  },
  methods: {
    handleSubmit: function Search(params) {
      axios.get('/s?q=' + this.movieTitle).then(response => {
        console.log("resp: ", response.data)
        if (response.data.status===0){
        this.info = genItem2(response.data)
        this.render = true
        this.status = false
        console.log("send: ", this.movieTitle)
       // console.log("rec: ", this.info)
        console.log("rec: ", this.info[0].film.title)
        } else{
        console.log("err")
        this.status = true
        this.render = false //fix await
        }
      })
     
    

    },

    handleChange(event) {
      this.movieTitle = event.target.value
     // console.log(this.movieTitle)
    }
  },
})
