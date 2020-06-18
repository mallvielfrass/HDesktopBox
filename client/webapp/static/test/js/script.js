
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    movieTitle: "",
    info2: [{film: {
      idKinopoisk: 470036 ,
      slogan: "«From the creators of Despicable Me»" ,
      title: "Лоракс" ,
      countries: " США" ,
      poster: "https://thumbs.bookdline.live/s/posters/thumbs/w220/loraks_film_2012_37834_0.jpeg" ,
      directors: " Кристоффер Боэ" ,
      short_story: "Так, много лет спустя, по вине человека, с планеты исчезли все деревья. Воздух теперь можно покупать в отдельных бутылочках, а внешний мир радует своей красотой. Одно желание - и у тебя под окном океан, второй - а там лес. Одно отличие от прошлой жизни, что все это пластиковое. Но многие дети, не видели настоящих деревьев, и потому уже не знают, чем этот мир может быть хуже. Теду всего двенадцать лет, но у него уже есть девушка. Одри, которую он так сильно любит, любит мечтать, а еще ее самое большое желание - увидеть живое дерево. Самое настоящее, которое видели еще ее родители. У Теда нет желания сказать, Одри нет, потому сначала, он спрашивает у бабушки, а потом отправляется к знакомому, который видел последнее дерево. И как оказывается, именно по его вине все это произошло, когда-то давно мужчине не послушал духа леса Лоракса и уничтожил все деревья. Теперь Теду предстоит посадить последнее дерево и найти духа, но это будет сложно, так как против этого выступают и власти и родители мальчика." ,
      year: 2012 ,
      original_title: "Dr. Seuss\' The Lorax"
      } }],
    info: {},

    render: false,
  },
  methods: {
    handleSubmit: function Search(params) {
      axios.get('/j?q=' + this.movieTitle).then(response => {
        this.info = genItem2(response.data)
        this.render = true
        console.log("send: ", this.movieTitle)
        console.log("rec: ", this.info)
        console.log("rec: ", this.info[0].film.title)
      })
     
    

    },

    handleChange(event) {
      this.movieTitle = event.target.value
      console.log(this.movieTitle)
    }
  },
})
