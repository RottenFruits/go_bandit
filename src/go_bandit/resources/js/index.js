var app6 = new Vue({
    el: '#app-6',
    data: {
        message: 'Hello!'
    }
})

var app_ = new Vue({ 
    el: '#app_',
    data:{
        selected:2,
        start_flag: false,
        treasures: [
            {state: 0, key: 0, link:"/resources/images/kaizoku_takarabako.png"}, 
            {state:0, key: 1, link:"/resources/images/kaizoku_takarabako.png"}
        ],
        arm_probs: [
            {prob: 0.3, key: 0}, 
            {prob: 0.3, key: 1}
        ]
    },
    methods:{
        start:function(){
            this.start_flag = true
            console.log(this.treasures)
            this.treasures.splice(0, 1)
            console.log("start")
        },
        stop:function(){
            this.start_flag = false
            this.treasures.push({state:0, key:2, link:"/resources/images/kaizoku_takara.png"});
            console.log("stop")            
        },
        clear:function(){
            console.log("clear")            
        },
        number_of_arms_select:function(){
            console.log(this.selected)
            this.selected = Number(this.selected)
        }
    }
})


var app = new Vue({
    el:"#app",　
    data:{　
      url:"http://localhost:8080/a",
      param:"{}",             
      result:"...."  
    },
    methods:{　  
      post: function(){
        config = {
          headers:{
            'X-Requested-With': 'XMLHttpRequest',
            'Content-Type':'application / x-www-form-urlencoded'
          },
          withCredentials:true,
        }
　　　　　 
        param = JSON.parse(this.param)

        axios.post(this.url,param,config)
        .then(function(res){
          app.result = res.data
          console.log(res)
        })
        .catch(function(res){
　　　　　　app.result = res.data
          console.log(res)
        })
      }
    }
  })