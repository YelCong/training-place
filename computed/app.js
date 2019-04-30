var vm = new Vue({
    el: '#app',
    data: {
        message: 'Hello',
        firstName: 'halo',
        lastName: 'World'
    },
    methods: {
        reversedMessage: function () {
            console.log('每有一处获取reversedMessage都可以看到我，我是没有缓存的');
            return this.message.split('').reverse().join('')
        },
        funcClick: function () {
            var str = 'halo ' + Math.ceil(Math.random() * 10);
            console.log(str);
            this.fullName = str;
        }
    },
    computed: {
        reversedMessage1: function () {
            // `this` 指向 vm 实例
            //@mark 计算属性是基于它们的响应式依赖进行缓存的
            console.log('我是不会多次输出的，因为我缓存了');
            return this.message.split('').reverse().join('');
        },
        cp_now:function () {
            console.log('同样的我也不会多次计算');
          return Date.now();
        },
        fullName:{
            get:function () {
                return this.firstName + '' + this.lastName
            },
            set:function (newValue) {
                var names = newValue.split(' ');
                this.firstName = names[0];
                this.lastName = names[names.length - 1];
            }
        }
    }
});