
var app= new Vue({
    el:'#app',
    data:{
        count:0,
        items: [
            { message: {text:"hi"} },
            { message: {text:"bye"} }
          ],
          info: {}, 
    },
    methods:{

        click: function(){
            this.count= this.count+1

            const [first,t,e] =genItem()
            console.log("data  returned", first,t,e);
        }
    },
    beforeMount(){
        this.items[1].message.text="world"
        this.count= this.count+1
        first = genItem()
       // console.log("data  returned", first.filminfo);
        this.info=first.filminfo[0]
        //console.log("data  returned", info);
     },
    




});
