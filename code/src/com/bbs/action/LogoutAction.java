package com.bbs.action;


/**
 * 
 * @author zhangyr
 * @version 1.0
 * 2021年3月23日下午10:15:56
 */
public class LogoutAction extends BaseAction {
	@Override
	public String execute() throws Exception {
		if (getSession().get("username")!=null){
			getSession().put("username",null);
			getSession().put("userId",null);
		}
		
		else if (getSession().get("adminname")!= null){
			getSession().put("adminname",null);
			getSession().put("adminid",null);
		}
		return SUCCESS;
	}

}
