    
window.onload=function(){
    // Step:3 conifg ECharts's path, link to echarts.js from current page.
    // Step:3 为模块加载器配置echarts的路径，从当前页面链接到echarts.js，定义所需图表路径
    require.config({
        paths: {
            echarts: '/static/js/echarts-2.2.6/build/dist'
        }
    });
    
    // Step:4 require echarts and use it in the callback.
    // Step:4 动态加载echarts然后在回调函数中开始使用，注意保持按需加载结构定义图表路径
    require(
        [
        'echarts',
        'echarts/chart/force',
        'echarts/chart/map'
        ],
        function (ec) {
            var myChart = ec.init(document.getElementById('relative-chart'));
            var nodes = [];
            var links = [];
            var constMaxDepth = 1;
            var constMaxChildren = 3;
            var constMinChildren = 3;
            var constMaxRadius =12;
            var constMinRadius = 11;

            function rangeRandom(min, max) {
                return Math.random() * (max - min) + min;
            }

            function createRandomNode(depth) {
                var node = {
                    name : 'NODE_' + nodes.length,
                    value : rangeRandom(constMinRadius, constMaxRadius),
        // Custom properties
        id : nodes.length,
        depth : depth,
        category : depth === constMaxDepth ? 0 : 1
    }
    nodes.push(node);

    return node;
}

function forceMockThreeData() {
    var depth = 0;
    var rootNode = {
        name : 'ROOT',
        value : rangeRandom(constMinRadius, constMaxRadius),
        // Custom properties
        id : 0,
        depth : 0,
        category : 2
    }
    nodes.push(rootNode);

    function mock(parentNode, depth) {
        var nChildren = Math.round(rangeRandom(constMinChildren, constMaxChildren));
        
        for (var i = 0; i < nChildren; i++) {
            var childNode = createRandomNode(depth);
            links.push({
                source : parentNode.id,
                target : childNode.id,
                weight : 1 
            });
            if (depth < constMaxDepth) {
                mock(childNode, depth + 1);
            }
        }
    }

    mock(rootNode, 0);
}

forceMockThreeData();
var option = {
    // title : {
    //     text: '你，Found，校园小伙伴',
    //     x:'right',
    //     y:'bottom'
    // },
    tooltip : {
        trigger: 'item',
        formatter: '{a} : {b}'
    },
    // toolbox: {
    //     show : true,
    //     feature : {
    //         restore : {show: true},
    //         magicType: {show: true, type: ['force', 'chord']},
    //         saveAsImage : {show: true}
    //     }
    // },
    legend: {
        show:false,
        x: 'left',
        data:['叶子节点','非叶子节点', '根节点']
    },
    series : [
    {
        type:'force',
        name : "Force tree",
        ribbonType: false,
        categories : [
        {
            name: '叶子节点'
        },
        {
            name: '非叶子节点'
        },
        {
            name: '根节点'
        }
        ],
        itemStyle: {
            normal: {
                label: {
                    show: false
                },
                nodeStyle : {
                    brushType : 'both',
                    borderColor : 'rgba(255,215,0,0.6)',
                    borderWidth : 1
                }
            }
        },
        minRadius : constMinRadius,
        maxRadius : constMaxRadius,
        coolDown: 0.995,
        steps: 10,
        nodes : nodes,
        links : links,
        steps: 1
    }
    ]
};

myChart.setOption(option);
}
);
}


