$(function(){
	$('.agree').on('click',function(){
		$('.step1').hide();
		$('.step2').fadeIn();
	})
	$('.disagree').on('click',function(){
		return false;
	})
	var checkEmail=false,
	checkUsername=false,
	checkPassword=false,
	checkpoint='<span class="checkpoint" style="font-size: 3px;color: red;">格式有误！</span>',
	countEmail=0,
	countUsername=0,
	countPassword=0;
	
	$('#email').on('blur',function(){
		countEmail++;
		textEmail.call(this,countEmail);
	})
	
	
	$('#username').on('blur',function(){
		countUsername++;
		textUsername.call(this,countUsername);
	})
	
	
	$('#password').on('blur',function(){
		countPassword++;
		textPassword.call(this,countPassword);
	})
	$('#registbtn').on('click',function(){
		if(checkEmail&&checkUsername&&checkPassword){
			console.log(checkEmail+checkUsername+checkPassword);
			$('#regist').submit();
			$('.step1,.step2').hide();
            $('.step3').slideDown();
		}else{
			alert('请将信息填写完整！');
		}
		
	})
	function textEmail(countEmail){
		var re=/[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/;
		var val= this.value;
		if(!re.test(val)){  //如果不合格
			if(countEmail!==1)return false;
			$(this).after(checkpoint);	
		}else{
			checkEmail=true;
			$(this).next('span').remove();
		}	
	}
	function textUsername(){
		var re=/^[a-zA-Z\u4e00-\u9fa5][a-zA-Z0-9_\u4E00-\u9FA5]{3,15}$/;
		var val=this.value;
		if(!re.test(val)){  //如果不合格
			if(countUsername!==1)return false;
			$(this).after(checkpoint);	
		}else{
			checkUsername=true;
			$(this).next('span').remove();
		}
	}
	function textPassword(){
		var re=/^[A-Za-z0-9_-]{5,15}$/;
		var val=this.value;
		if(!re.test(val)){  //如果不合格
			if(countPassword!==1)return false;
			$(this).after(checkpoint);	
		}else{
			checkPassword=true;
			$(this).next('span').remove();
		}
	}
})
