<!-- 因为 AJAX 库和通用工具的生态已经相当丰富，Vue 核心代码没有重复 -->
<!-- 提供这些功能以保持精简。这也可以让你自由选择自己更熟悉的工具。 -->
var watchExampleVM = new Vue({
        el: '#watch-example',
        data: {
            question: '',
            answer: 'I cannot give you an answer until you ask a question!'
        },
        watch: {
            // 如果 `question` 发生改变，这个函数就会运行
            question: function (newQuestion, oldQuestion) {
                this.answer = 'Waiting for you to stop typing...';
                //@mark 这个地方不直接调用getAnswer只是为了避免频繁请求。
                this.debouncedGetAnswer();
            }
        },
        created: function () {
            console.log('vue实例生成时执行的。');
            // `_.debounce` 是一个通过 Lodash 限制操作频率的函数。
            // 在这个例子中，我们希望限制访问 yesno.wtf/api 的频率
            // AJAX 请求直到用户输入完毕才会发出。想要了解更多关于
            // `_.debounce` 函数 (及其近亲 `_.throttle`) 的知识，
            // 请参考：https://lodash.com/docs#debounce
            this.debouncedGetAnswer = _.debounce(this.getAnswer, 500)
        },
        methods: {
            getAnswer: function () {
                console.log(123);
                if (this.question.indexOf('?') === -1) {
                    this.answer = 'Questions usually contain a question mark. ;-)'
                    return
                }
                this.answer = 'Thinking...'
                var vm = this
                axios.get('https://yesno.wtf/api')
                    .then(function (response) {
                        vm.answer = _.capitalize(response.data.answer)
                    })
                    .catch(function (error) {
                        vm.answer = 'Error! Could not reach the API. ' + error
                    })
            }
        }
    })