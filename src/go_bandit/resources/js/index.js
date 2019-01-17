var app_ = new Vue({
    el: '#app_',
    data: {
        selected: 2,
        auto_flag: false,
        auto_mode_flag: false,
        timer_id: null,
        trials: 0,
        cumulative_rewards: 0,
        treasures: [
            { state: 0, key: 0, link: "/resources/images/kaizoku_takarabako.png", visible: true },
            { state: 0, key: 1, link: "/resources/images/kaizoku_takarabako.png", visible: true },
            { state: 0, key: 2, link: "/resources/images/kaizoku_takarabako.png", visible: false },
            { state: 0, key: 3, link: "/resources/images/kaizoku_takarabako.png", visible: false },
            { state: 0, key: 4, link: "/resources/images/kaizoku_takarabako.png", visible: false }
        ],
        arm_parameters: [
            { prob: 0.3, key: 0, visible: true, counts: 0, values: 0, arm_rewards: 0, cvr: 0 },
            { prob: 0.3, key: 1, visible: true, counts: 0, values: 0, arm_rewards: 0, cvr: 0 },
            { prob: 0.3, key: 2, visible: false, counts: 0, values: 0, arm_rewards: 0, cvr: 0 },
            { prob: 0.3, key: 3, visible: false, counts: 0, values: 0, arm_rewards: 0, cvr: 0 },
            { prob: 0.3, key: 4, visible: false, counts: 0, values: 0, arm_rewards: 0, cvr: 0 }
        ],
        bandit: 
            {
                algorithm: "EG",
                epsilon: 0.8,
                n: 2,
                counts: [0, 0, 0, 0, 0],
                values: [0, 0, 0, 0, 0],
                arm_rewards: [0, 0, 0, 0, 0],
                chosen_arms: [], 
                rewards: [], 
                cumulative_rewards: []
            }
        
    },
    methods: {
        start: function () {
            console.log("start")
            //image setting
            this.image_initialize()

            //prameter cast
            this.bandit["n"] = Number(this.selected)
            this.bandit.epsilon = parseFloat(this.bandit.epsilon)
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

            url = "http://localhost:8080/oneshot"

            axios.post(url, {
                bandit: this.bandit,
                arm_parameters: this.arm_parameters
            },
                config)
                .then(function (res) {
                    //update results
                    for (i = 0; i < res.data.bandit.counts.length; i++) {
                        app_.bandit.counts[i] = res.data.bandit.counts[i]
                        app_.bandit.values[i] = res.data.bandit.values[i]
                        app_.bandit.arm_rewards[i] = res.data.bandit.arm_rewards[i]
                        //update display
                        app_.arm_parameters[i].counts = res.data.bandit.counts[i]
                        app_.arm_parameters[i].arm_rewards = res.data.bandit.arm_rewards[i]
                        app_.arm_parameters[i].values = res.data.bandit.values[i]
                        app_.arm_parameters[i].cvr = app_.arm_parameters[i].arm_rewards / app_.arm_parameters[i].counts
                    }
                    n_counts = res.data.bandit.chosen_arms.length
                    app_.bandit.chosen_arms.push(res.data.bandit.chosen_arms[n_counts - 1])
                    app_.bandit.rewards.push(res.data.bandit.rewards[n_counts - 1])
                    app_.bandit.cumulative_rewards.push(res.data.bandit.cumulative_rewards[n_counts - 1])
                    app_.trials = n_counts
                    app_.cumulative_rewards = res.data.bandit.cumulative_rewards[n_counts - 1]

                    //update image
                    chosen_treasure = res.data.bandit.chosen_arms[n_counts - 1]
                    if (res.data.bandit.rewards[n_counts - 1] == 1) {
                        app_.treasures[chosen_treasure].link = "/resources/images/kaizoku_takara.png"
                    } else {
                        app_.treasures[chosen_treasure].link = "/resources/images/character_game_mimic.png"
                    }
                    setTimeout(app_.image_initialize, 100)

                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        auto: function () {
            this.auto_flag = true
            this.start()
            this.timer_id = setInterval(this.start, 300)
            console.log("auto")
        },
        stop: function () {
            this.auto_flag = false
            clearInterval(this.timer_id)
            console.log("stop")
        },
        clear: function () {
            console.log("clear")
            for (i = 0; i < this.arm_parameters.length; i++) {
                //update display
                this.arm_parameters[i].counts = 0
                this.arm_parameters[i].arm_rewards = 0
                this.arm_parameters[i].values = 0
                this.arm_parameters[i].cvr = 0
            }
            this.trials = 0
            this.cumulative_rewards = 0
            this.bandit.counts = [0, 0, 0, 0, 0]
            this.bandit.values = [0, 0, 0, 0, 0]
            this.bandit.arm_rewards = [0, 0, 0, 0, 0]
            this.bandit.chosen_arms = []
            this.bandit.rewards = []
            this.bandit.cumulative_rewards = []
        },
        auto_mode_check: function () {
            if (this.auto_mode_flag == false) {
                this.stop()
            }
        },
        image_initialize: function () {
            this.treasures[0].link = "/resources/images/kaizoku_takarabako.png"
            this.treasures[1].link = "/resources/images/kaizoku_takarabako.png"
            this.treasures[2].link = "/resources/images/kaizoku_takarabako.png"
            this.treasures[3].link = "/resources/images/kaizoku_takarabako.png"
            this.treasures[4].link = "/resources/images/kaizoku_takarabako.png"
        },
        number_of_arms_select: function () {
            this.selected = Number(this.selected)
            for (i = 0; i < this.arm_parameters.length; i++) {
                if (i < this.selected) {
                    this.arm_parameters[i]['visible'] = true
                    this.treasures[i]['visible'] = true
                } else {
                    this.arm_parameters[i]['visible'] = false
                    this.treasures[i]['visible'] = false
                }
            }
        }

    }
})

var graph = new Vue({
    el: '#graph',
    data: {
        data1: 40,
        data2: 80
    }
})

Vue.component('bar', {
    mixins: [VueChartJs.Bar, VueChartJs.mixins.reactiveData],
    data: function () {
        return {
            options: {
                scales: {
                    yAxes: [
                        {
                            ticks: {
                                min: 0,
                                max: 100,
                            }
                        },
                    ]
                },
                responsive: false,
                data1: {
                    type: Number,
                    required: true,
                },
                data2: {
                    type: Number,
                    required: true,
                }
            },
        }
    },
    props: {
        data1: {
            type: Number,
            required: true,
        },
        data2: {
            type: Number,
            required: true,
        }
    },
    watch: {
        data1: function () {
            this.updateChartData()
        },
        data2: function () {
            this.updateChartData()
        }
    },
    methods: {
        updateChartData() {
            const newChartData = Object.assign({}, this.chartData)
            newChartData.datasets[0].data = [this.data1]
            newChartData.datasets[1].data = [this.data2]
            this.chartData = newChartData
        },
    },
    mounted: function () {
        this.chartData = {
            datasets: [
                {
                    label: "data1",
                    data: [this.data1],
                },
                {
                    label: "data2",
                    data: [this.data2],
                },
            ],
        }
    }
})

new Vue({ el: '#graph' })
