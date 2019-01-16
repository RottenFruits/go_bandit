var app_ = new Vue({
    el: '#app_',
    data: {
        selected: 2,
        auto_flag: false,
        auto_mode_flag: false,
        timer_id: null,
        treasures: [
            { state: 0, key: 0, link: "/resources/images/kaizoku_takarabako.png", visible: true },
            { state: 0, key: 1, link: "/resources/images/kaizoku_takarabako.png", visible: true },
            { state: 0, key: 2, link: "/resources/images/kaizoku_takarabako.png", visible: false },
            { state: 0, key: 3, link: "/resources/images/kaizoku_takarabako.png", visible: false },
            { state: 0, key: 4, link: "/resources/images/kaizoku_takarabako.png", visible: false }
        ],
        arm_parameters: [
            { prob: 0.3, key: 0, visible: true },
            { prob: 0.3, key: 1, visible: true },
            { prob: 0.3, key: 2, visible: false },
            { prob: 0.3, key: 3, visible: false },
            { prob: 0.3, key: 4, visible: false }
        ],
        bandit_results: [
            { chosen_arms: [], rewards: [], cumulative_rewards: [] }
        ],
        bandit: [{
            algorithm: "EG",
            epsilon: 0.8,
            n: 2,
            counts: [0, 0, 0, 0, 0],
            values: [0, 0, 0, 0, 0]
        }
        ]

    },
    methods: {
        start: function () {
            console.log("start")
            //prameter cast
            this.bandit[0]["n"] = Number(this.selected)
            for (i = 0; i < this.arm_parameters.length; i++) {
                this.arm_parameters[i]["prob"] = parseFloat(this.arm_parameters[i]["prob"])
            }

            config = {
                headers: {
                    'X-Requested-With': 'XMLHttpRequest',
                    'Content-Type': 'application/json'
                },
                withCredentials: true,
            }

            url = "http://localhost:8080/a"

            axios.post(url, {
                bandit: this.bandit,
                arm_parameters: this.arm_parameters,
                bandit_results: this.bandit_results
            },
                config)
                .then(function (res) {
                    app.result = res.data

                    //update results
                    for (i = 0; i < res.data.bandit[0].Counts.length; i++) {
                        app_.bandit[0].counts[i] = res.data.bandit[0].Counts[i]
                        app_.bandit[0].values[i] = res.data.bandit[0].Values[i]
                    }
                    app_.bandit_results[0].chosen_arms = res.data.bandit_results[0].Chosen_arms
                    app_.bandit_results[0].rewards = res.data.bandit_results[0].Rewards
                    app_.bandit_results[0].cumulative_rewards = res.data.bandit_results[0].Cumulative_rewards


                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        auto: function () {
            this.auto_flag = true
            this.timer_id = setInterval(this.start, 1000)
            console.log("auto")
        },
        stop: function () {
            this.auto_flag = false
            clearInterval(this.timer_id)
            console.log("stop")
        },
        clear: function () {
            console.log("clear")
            app_.bandit[0].counts = [0, 0, 0, 0, 0]
            app_.bandit[0].values = [0, 0, 0, 0, 0]
            app_.bandit_results[0].chosen_arms = []
            app_.bandit_results[0].rewards = []
            app_.bandit_results[0].cumulative_rewards = []
        },
        auto_mode_check: function(){
            if(this.auto_mode_flag == false){
                this.stop()
            }
        },
        number_of_arms_select: function () {
            this.selected = Number(this.selected)

            if (this.selected == 2) {
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
            } else if (this.selected == 3) {
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
            } else if (this.selected == 4) {
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
            } else if (this.selected == 5) {
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