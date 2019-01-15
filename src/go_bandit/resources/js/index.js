var app6 = new Vue({
    el: '#app-6',
    data: {
        message: 'Hello!'
    }
})

Vue.component('arm-prob', {
    data: function () {
      return {
        prob: 0.3,
        key:null
      }
    },
    template: '<input type="text" value=0.0 size="15"></input>'
})

Vue.component('treasure', {
    data: function () {
      return {
        state: 0,
        key:null
      }
    },
    template: '<img src="/resources/images/kaizoku_takarabako.png" width="100"></img>'
})


var app_ = new Vue({ 
    el: '#app_',
    data:{
        selected:2,
        start_flag: false,
        treasures: [
            {state: 0, key: 0}, 
            {state:0, key: 1}
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
            this.treasures.push({state:0, key:2});
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