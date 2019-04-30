var vm = new Vue({
    el: '#app',
    data: {
        isActive: true,
        hasError: false,
        classObject: {
            'active': true,
            'hasError': true,
            'isP': false,
            'isC': false
        }
    },
    methods: {
        funcClick: function () {
            this.classObject.isC = !this.classObject.isC;
        },
    }
});