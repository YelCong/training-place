var app = new Vue({
    el: '#app',
    data: {
        isA: false,
        isB: true,
        type: 'A',
        parentMessage: 'Parent',
        items: [
            { message: 'Foo' },
            { message: 'Bar' }
        ],
        object: {
            title: 'How to do lists in Vue',
            author: 'Jane Doe',
            publishedAt: '2016-04-10'
        }
    }
});