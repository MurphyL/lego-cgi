package wecom

// 提醒群中的所有人成员
func AtAll() messageOption {
	return at("mentioned_list", []string{"@all"})
}

// 通过userId提醒群中的指定成员(@某个成员)
func AtUserIds(userIds ...string) messageOption {
	return at("mentioned_list", userIds)
}

// 通过手机号提醒群中的指定成员(@某个成员)
func AtMobiles(mobiles ...string) messageOption {
	return at("mentioned_mobile_list", mobiles)
}

// 通过手机号提醒群中的指定成员，@all表示提醒所有人
func at(key string, refs []string) messageOption {
	return func(msg messagePack) {
		msg[key] = refs
	}
}
