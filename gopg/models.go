package main

type Company struct {
	ID    uint64
	Users []*User `pg:"many2many:company_users"`
}

type User struct {
	ID        uint64
	Companies []*Company `pg:"many2many:company_users"`
}

type CompanyUser struct {
	ID        uint64
	CompanyID uint64
	Company   *Company
	UserID    uint64
	User      *User
}
