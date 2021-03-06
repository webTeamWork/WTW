package com.bbs.serviceImpl;

import com.bbs.dao.FollowcardDao;
import com.bbs.hibernate.factory.BaseHibernateDAO;
import com.bbs.model.Followcard;
import com.bbs.service.FollowcardBiz;
/**
 * 
 * @author zhangyr
 * @version 1.0
 * 2021年3月23日上午11:32:41
 */
public class FollowcardBizImpl implements FollowcardBiz{
	
	private FollowcardDao followcardDao;
	
	public void setFollowcardDao(FollowcardDao followcardDao) {
		this.followcardDao = followcardDao;
	}

	/* (non-Javadoc)
	 * @see com.bbs.bizImpl.FollowcardBiz#addReply(com.bbs.model.Followcard)
	 */
	@Override
	public void addReply(Followcard followcard){
		followcardDao.save(followcard);
	}

}
