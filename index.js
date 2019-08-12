var Main = {
    data() {
        let d = idxData();
        console.log(d.data);
        return {
            input: '',
            curPage: d.page,
            pageSize: d.limit,
            count: d.count,
            tableData: d.data.list,
            dialogTableVisible: false,
            dialogFormVisible: false,
            form: {
                name: '',
                region: '',
                date1: '',
                date2: '',
                delivery: false,
                type: [],
                resource: '',
                desc: ''
            },
            formLabelWidth: '120px'
        };
    },
    mouted(){
        this.pageInit(1,20)
    },
    methods: {
        pageInit:function (paeg,pageSize) {
            let d = idxData();
            this.curPage = d.page;
            this.pageSize = d.limit;
            this.count = d.count;
            this.tableData = d.data.list;
        },
        groupDel:function(row) {
            let vm = this;
            console.log(this);
            console.log(row);
            this.$confirm('此操作将永久删除该文件', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(function () {
                console.log('这里做ajax请求删除接口');
                vm.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                vm.pageInit(vm.curPage,vm.pageSize);
            }).catch(
                function () {
                    vm.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                }
            );
        }
    }
};
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')

function idxData(page,pageSize) {
    console.log('1');
    var str = '{"code":0,"message":"操作成功","data":{"page":1,"limit":20,"count":16,"list":[{"id":"58e96206452e987ef6a5ba5a","creator_id":"5a9e0b28e9e85220f9f8fdb9","name":"叶赐红测试","intro":"111","member_num":0,"created":1564726732,"updated":1564726732,"creator_name":"叶赐红"},{"id":"f95060aa449ca0b25ec40630","creator_id":"5b5edbb68ba41201a6000032","name":"sy测试组","intro":"","member_num":0,"created":1563185993,"updated":1563185993,"creator_name":"蔡奕勇"},{"id":"6e682cd14e2dacd7ca81b5a0","creator_id":"5b5edbb68ba41201a6000032","name":"好游快爆-运维","intro":"","member_num":0,"created":1557981070,"updated":1557981070,"creator_name":"蔡奕勇"},{"id":"fa610ff945b0bf524225a57a","creator_id":"58ef44b0bdb6ac1528000001","name":"4399运维","intro":"4399运维","member_num":0,"created":1553678920,"updated":1553679209,"creator_name":"陈兴洪"},{"id":"94a55e7245d5aa27d3b75261","creator_id":"58ef44b0bdb6ac1528000001","name":"您说得对","intro":"东哥说得对","member_num":0,"created":1545114848,"updated":1551683606,"creator_name":"陈兴洪"},{"id":"d3e373b34fa9a71f88db363a","creator_id":"5a1fce6d10389e018cf659db","name":"云笔记测试","intro":"","member_num":0,"created":1545113350,"updated":1545113350,"creator_name":"欧强"},{"id":"db29e8a84110b19302537c60","creator_id":"58ef44b0bdb6ac1528000001","name":"服务端","intro":"服务端开发","member_num":0,"created":1542953581,"updated":1542954269,"creator_name":"陈兴洪"},{"id":"f1dc14134857a5d5a67ca3fb","creator_id":"58ef44b0bdb6ac1528000001","name":"游拍服务端","intro":"游拍服务端开发","member_num":0,"created":1542953611,"updated":1542954255,"creator_name":"陈兴洪"},{"id":"e1aca61a46489a2e3e4437de","creator_id":"5a1fce6d10389e018cf659db","name":"游戏盒服务端","intro":"游戏盒服务端开发","member_num":0,"created":1542953552,"updated":1542954251,"creator_name":"欧强"},{"id":"ee79a8f041e3b988bec2440a","creator_id":"5a1fce6d10389e018cf659db","name":"活动项目组","intro":"","member_num":0,"created":1541389136,"updated":1541389136,"creator_name":"欧强"},{"id":"2b7bfdec4d929f28c58708c4","creator_id":"5a1fce6d10389e018cf659db","name":"喵叔小组","intro":"","member_num":0,"created":1541135894,"updated":1541135894,"creator_name":"欧强"},{"id":"cf539f6f4512889978056df8","creator_id":"5a1fce6d10389e018cf659db","name":"游戏吧社区","intro":"","member_num":0,"created":1541135329,"updated":1541135329,"creator_name":"欧强"},{"id":"63dc763a4788b72600587231","creator_id":"5a1fce6d10389e018cf659db","name":"运营","intro":"","member_num":0,"created":1541135254,"updated":1541135254,"creator_name":"欧强"},{"id":"606820174b288c4f29d9402b","creator_id":"5a1fce6d10389e018cf659db","name":"游戏盒","intro":"","member_num":0,"created":1541135082,"updated":1541135082,"creator_name":"欧强"},{"id":"96be94bd4e908b666768afd7","creator_id":"5a1fce6d10389e018cf659db","name":"游拍","intro":"","member_num":0,"created":1541135082,"updated":1541135082,"creator_name":"欧强"},{"id":"6d936bba46b5b067c2e43d66","creator_id":"5a1fce6d10389e018cf659db","name":"公开群","intro":"所有成员自动加入公开群","member_num":0,"created":1540798963,"updated":1540798963,"creator_name":"欧强"}]}}';
    return eval('(' + str + ')');
}

function members() {
    var str = '{"code":0,"message":"操作成功","data":{"list":[{"id":"5d2da6d20463590781000422","name":"黄敬平","email":"huangjingping@4399inc.com","created":1563272915,"updated":1563272915,"type":1},{"id":"5d2da7d50463590781000427","name":"黄剑华","email":"huangjianhua@4399inc.com","created":1563273173,"updated":1563273173,"type":1},{"id":"5d2e8474046359078100044c","name":"蓝锟","email":"lankun@4399inc.com","created":1563329652,"updated":1563329652,"type":1}]}}';
    return eval('(' + str + ')');
}

function tree() {
    var str = '{"code":0,"message":"操作成功","data":{"list":[{"id":"421342b04778bca0c623450e","text":"4399运维","pid":"","children":[{"id":"5d43af6537bd766f9f000000","text":"文件存档","pid":"421342b04778bca0c623450e"},{"id":"5cee4e3f5d104317ba00000b","text":"总结和规划","pid":"421342b04778bca0c623450e","children":[{"id":"5d1c654c5d10430ba4000000","text":"安全总结","pid":"5cee4e3f5d104317ba00000b"},{"id":"5d0c48c65d10435a63000002","text":"年度总结","pid":"5cee4e3f5d104317ba00000b"},{"id":"5d0c48bb5d10435a63000001","text":"记录文档","pid":"5cee4e3f5d104317ba00000b"},{"id":"5d0c48b35d10435a63000000","text":"会议记录","pid":"5cee4e3f5d104317ba00000b"}]},{"id":"5cd3dfdb3ba98d3ae1000000","text":"运维平台","pid":"421342b04778bca0c623450e","children":[{"id":"5cee4f0e5d104317ba000011","text":"cloud.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5cee4eee5d104317ba00000d","text":"jenkins.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5cee4ee15d104317ba00000c","text":"gitlab.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5ce4be213ba98d0fcf000000","text":"rzxt.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5cd92b023ba98d5cb000000e","text":"zabbix.5054399.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5cd928693ba98d5cb000000b","text":"job.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000"},{"id":"5cd9284f3ba98d5cb000000a","text":"devops.4399doc.com","pid":"5cd3dfdb3ba98d3ae1000000","children":[{"id":"5d03006a5d10430c21000001","text":"架构文档","pid":"5cd9284f3ba98d5cb000000a"},{"id":"5d03003e5d10430c21000000","text":"帮助文档","pid":"5cd9284f3ba98d5cb000000a"}]}]},{"id":"2e75f02e4790bfd5e33f2d3c","text":"流程制度","pid":"421342b04778bca0c623450e"},{"id":"5c9c85dab021c12532000000","text":"项目文档","pid":"421342b04778bca0c623450e","children":[{"id":"5d2c528cafd6b346b5000004","text":"原创平台","pid":"5c9c85dab021c12532000000"},{"id":"5d11e8aac2f9526791000000","text":"造梦西游OL","pid":"5c9c85dab021c12532000000"},{"id":"5cf723a3afd6b36a9b000000","text":"孩子站(猫小帅)","pid":"5c9c85dab021c12532000000"},{"id":"5cf62d5b5d10433de6000013","text":"桌趣","pid":"5c9c85dab021c12532000000"},{"id":"5cf62d2d5d10433de6000011","text":"用户中心","pid":"5c9c85dab021c12532000000"},{"id":"5cf62d0b5d10433de600000f","text":"小小突击队","pid":"5c9c85dab021c12532000000"},{"id":"5cf62ccd5d10433de600000c","text":"西游之超神","pid":"5c9c85dab021c12532000000"},{"id":"5cf62c875d10433de600000a","text":"热血赛车","pid":"5c9c85dab021c12532000000"},{"id":"5cf62c535d10433de6000008","text":"全站统计","pid":"5c9c85dab021c12532000000"},{"id":"5cf62be95d10433de6000006","text":"趣核游戏平台","pid":"5c9c85dab021c12532000000"},{"id":"5cf62bbe5d10433de6000004","text":"漫画盒","pid":"5c9c85dab021c12532000000"},{"id":"5cf629e95d10433de6000002","text":"独代平台","pid":"5c9c85dab021c12532000000"},{"id":"5cf629b55d10433de6000000","text":"斗地主","pid":"5c9c85dab021c12532000000"},{"id":"5ce218d53ba98d3aaa000000","text":"游戏盒-活动","pid":"5c9c85dab021c12532000000"},{"id":"5cd91f983ba98d5cb0000007","text":"SDK联运m.4399api.com","pid":"5c9c85dab021c12532000000"},{"id":"5cd8dfb13ba98d5cb0000001","text":"游拍","pid":"5c9c85dab021c12532000000"},{"id":"5cd385f53ba98d4731000000","text":"好游快爆","pid":"5c9c85dab021c12532000000"},{"id":"5cc5676c3ba98d2d7b000003","text":"生死狙击","pid":"5c9c85dab021c12532000000"},{"id":"5cc567633ba98d2d7b000002","text":"游戏盒","pid":"5c9c85dab021c12532000000"},{"id":"5c9c85efb021c12532000001","text":"游戏吧","pid":"5c9c85dab021c12532000000"}]},{"id":"60c6eb3648a88903a7d18563","text":"技术分享","pid":"421342b04778bca0c623450e","children":[{"id":"5d2c48a7afd6b346b5000000","text":"技术文档","pid":"60c6eb3648a88903a7d18563"},{"id":"5d0c4af65d10435a63000008","text":"实战操作","pid":"60c6eb3648a88903a7d18563"},{"id":"5d0c531f5d10435a6300000a","text":"原理研究","pid":"60c6eb3648a88903a7d18563"},{"id":"5d0c4a9b5d10435a63000006","text":"故障总结","pid":"60c6eb3648a88903a7d18563"}]}]}]}}';
    return eval('(' + str + ')');
}

function edit() {
    var str = '{"code":0,"message":"操作成功","data":{}}';
}

function f() {

}