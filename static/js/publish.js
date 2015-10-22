$(function(){
	var str='<tr>\
                <td><input id="number" type="number"></td>\
                    <td>\
                        <select id="classify">\
                            <option>web前端</option>\
                            <option>C++</option>\
                            <option>java后台</option>\
                            <option>php</option>\
                            <option>高等数学</option>\
                        </select>\
                    </td>\
                    <td>\
                        <input type="text" >\
                    </td>\
            </tr>';
    $("#add").click(function(){
         $(".member-tab tbody").append(str);
    });
    
    
    $(".publish-btn").click(function(){
        createData();
    });

});

function createData(){
    var data={
            title:$("#pub-title")[0].value,
            content:$("#pub-detail")[0].value,
            member:[],
            deadline:$("#deadline")[0].value,
          contactWay:{
                    phoneNumber:$("#phone")[0].value,
                    QQNumber:$("#qq")[0].value,
                    email:$("#email")[0].value,
                    wechatNumber:$("#we-chat")[0].value
            }

        },
        
        oTr=$(".member-tab tbody tr");

    
   
   
    for(var i=0;i<oTr.length;i++){
        var oTd=oTr[i].getElementsByTagName('td'),
            singleMember={
               number:0,
               skill:"",
               ps:""
            };
        

        singleMember["number"]=oTd[0].getElementsByTagName("input")[0].value;
        singleMember["skill"]=oTd[1].getElementsByTagName("select")[0].value;
        singleMember["ps"]=oTd[2].getElementsByTagName("input")[0].value;

        data.member.push(singleMember);
    }

    $.ajax({
            type:'post',
            url:'/api/froms/changs/publish',
            dataType:'json',
            data:{
                'data':data

            },
            async:false,
            success:function(data){
                
            }
        });

    
}