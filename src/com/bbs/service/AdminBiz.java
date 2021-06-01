package com.bbs.service;

import com.bbs.model.Admin;
import com.bbs.model.User;
/**
 * 
 * @author zhangyr
 * @version 1.0
 * 2021年3月23日上午11:35:37
 */
public interface AdminBiz {

	/**
	 * 用户登陆
	 * @param user 用户对象
	 * @return 
	 */
	public abstract int login(String username, String password);

	public abstract void updateAdmin(Admin admin);

	public abstract Admin getAdminById(Integer adminId);


	public abstract int getAdminIdByEmail(String email);

	public abstract int isExist(Admin admin);

	public abstract int getAdminIdByUsername(String username);

}