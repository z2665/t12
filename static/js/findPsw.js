$(function(){
    $(".step1 .submit-btn").click(function(ev){
        var accountInfo={
            username:$(".step1 #username").val(),
            email:$(".step1 #username-email").val()
        },
        jsonAccountInfo=JSON.stringify(accountInfo); 

        // $(".step1 .error-tip").removeClass("hide");
        // console.log(jsonAccountInfo);

        // $(".step1").fadeOut(200,function(){
        //                 $(".step2").fadeIn(200);
        //             });
        ev.preventDefault();

        $.ajax({
            type:'post',
            url:"/api/users/forget",
            dataType:'json',
            data:{
                'data':jsonAccountInfo
            },
            async:false,
            success:function(data){
                
                if(data.data){
                    $(".step1").fadeOut(200,function(){
                        $(".step2").fadeIn(200);
                    });
                }else{
                    $(".step1 .error-tip").removeClass("hide");
                }
            }
        });

    });

    $(".step2 .submit-btn").click(function(){
        var updata={
            "id":$(".step2 #guard-code").val(),
            "password":$(".step2 #psd").val()
        },
            jsonUpdata=JSON.stringify(updata);

        console.log(jsonUpdata);

        $.ajax({
            type:'post',
            url:"/api/users/forgetresult",
            dataType:'json',
            data:{
                'data':jsonUpdata
            },
            async:false,
            success:function(data){
                
                if(data.data){
                    alert("更新信息成功");
                }else{
                    alert("更新信息失败");
                }
            }
        });

    });
})
