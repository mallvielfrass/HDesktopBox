
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    movieTitle: "",
    info: [],
    info2: {},
    render: false,
  },
  methods: {
    handleSubmit: function Search(params) {
      params.preventDefault()
      console.log("send: ", this.movieTitle)
      var xhr = new XMLHttpRequest();
      xhr.open('GET', '/s?q=' + this.movieTitle, false);
      xhr.send();
      var jsonResponse = JSON.parse(xhr.responseText);
      this.info = genItem2(jsonResponse)
      this.render = true
      console.log("send: ", this.movieTitle)
      console.log("rec: ", this.info)
      console.log("rec: ", this.info[0].film.title)

    },

    handleChange(event) {
      this.movieTitle = event.target.value
      console.log(this.movieTitle)
    }
  },
})
