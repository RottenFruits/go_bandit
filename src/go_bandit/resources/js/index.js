var app_ = new Vue({ 
    el: '#app_',
    data:{
        selected:2,
        start_flag: false,
        treasures: [
            {state: 0, key: 0, link:"/resources/images/kaizoku_takarabako.png", visible:true}, 
            {state:0, key: 1, link:"/resources/images/kaizoku_takarabako.png", visible:true},
            {state:0, key: 2, link:"/resources/images/kaizoku_takarabako.png", visible:false},
            {state:0, key: 3, link:"/resources/images/kaizoku_takarabako.png", visible:false},
            {state:0, key: 4, link:"/resources/images/kaizoku_takarabako.png", visible:false}
        ],
        arm_parameters: [
            {prob: 0.3, key: 0, visible:true}, 
            {prob: 0.3, key: 1, visible:true},
            {prob: 0.3, key: 2, visible:false},
            {prob: 0.3, key: 3, visible:false},
            {prob: 0.3, key: 4, visible:false}
        ]
    },
    methods:{
        start:function(){
            console.log("start")
            this.start_flag = true

            n_arms = Number(this.selected)
            config = {
                headers:{
                  'X-Requested-With': 'XMLHttpRequest',
                  'Content-Type':'application/json'
                },
                withCredentials:true,
              }
              
              url = "http://localhost:8080/a"
      
                axios.post(url,{
                    n_arms:n_arms, 
                    arm_parameters:this.arm_parameters
                  },
                  config)
                .then(function(res){
                  app.result = res.data
                  console.log(res)
                })
                .catch(function(error){
                  console.log(error)
                })  

            
        },
        stop:function(){
            this.start_flag = false
            console.log("stop")            
        },
        clear:function(){
            console.log("clear")            
        },
        number_of_arms_select:function(){
            this.selected = Number(this.selected)

            if(this.selected == 2){
                this.arm_parameters[0]['visible'] = true
                this.arm_parameters[1]['visible'] = true
                this.arm_parameters[2]['visible'] = false
                this.arm_parameters[3]['visible'] = false
                this.arm_parameters[4]['visible'] = false
                this.treasures[0]['visible'] = true
                this.treasures[1]['visible'] = true
                this.treasures[2]['visible'] = false
                this.treasures[3]['visible'] = false
                this.treasures[4]['visible'] = false
            }else if(this.selected == 3){
                this.arm_parameters[0]['visible'] = true
                this.arm_parameters[1]['visible'] = true
                this.arm_parameters[2]['visible'] = true
                this.arm_parameters[3]['visible'] = false
                this.arm_parameters[4]['visible'] = false
                this.treasures[0]['visible'] = true
                this.treasures[1]['visible'] = true
                this.treasures[2]['visible'] = true
                this.treasures[3]['visible'] = false
                this.treasures[4]['visible'] = false
            }else if(this.selected == 4){
                this.arm_parameters[0]['visible'] = true
                this.arm_parameters[1]['visible'] = true
                this.arm_parameters[2]['visible'] = true
                this.arm_parameters[3]['visible'] = true
                this.arm_parameters[4]['visible'] = false
                this.treasures[0]['visible'] = true
                this.treasures[1]['visible'] = true
                this.treasures[2]['visible'] = true
                this.treasures[3]['visible'] = true
                this.treasures[4]['visible'] = false
            }else if(this.selected == 5){
                this.arm_parameters[0]['visible'] = true
                this.arm_parameters[1]['visible'] = true
                this.arm_parameters[2]['visible'] = true
                this.arm_parameters[3]['visible'] = true
                this.arm_parameters[4]['visible'] = true
                this.treasures[0]['visible'] = true
                this.treasures[1]['visible'] = true
                this.treasures[2]['visible'] = true
                this.treasures[3]['visible'] = true
                this.treasures[4]['visible'] = true
            }
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