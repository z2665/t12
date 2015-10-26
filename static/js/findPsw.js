$(function(){
    $(".step1 .submit-btn").click(function(){
        var accountInfo={
        	phoneor:$(".step1 #username").val(),
        	vCode:$(".step1 #v-code").val()
        },
        jsonAccountInfo=JSON.stringify(accountInfo); 
        $(".step1 .error-tip").removeClass("hide");
        console.log(jsonAccountInfo);

        $(".step1").fadeOut(200,function(){
    		            $(".step2").fadeIn(200);
    	            });
        $.ajax({
            type:'post',
            url:"",
            dataType:'json',
            data:{
            	'data':jsonAccountInfo
            },
            async:false
            success:function(data){
                
                if(data.status){
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
        	username:$(".step2 #username").val(),
        	newPsw:$(".step2 #psd").val(),
        	vCode:$(".step2 .send-code").val()
        },
            jsonUpdata=JSON.stringify(updata);
        
        $.ajax({
            type:'post',
            url:"",
            dataType:'json',
            data:{
            	'data':jsonUpdata
            },
            async:false
            success:function(data){
                
                if(data.status){
                    alert("更新信息成功");
                }else{
                	alert("更新信息失败");
                }
            }
        });



    });


})