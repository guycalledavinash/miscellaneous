from typing import Any, Generic, TypeVar
from sqlalchemy.orm import Session
T = TypeVar("T")
class Repository(Generic[T]):
    def __init__(self, db: Session, model: type[T]): self.db, self.model = db, model
    def list(self, skip=0, limit=50): return self.db.query(self.model).offset(skip).limit(limit).all()
    def get(self, item_id: int): return self.db.get(self.model, item_id)
    def create(self, data: dict[str, Any]):
        obj = self.model(**data); self.db.add(obj); self.db.commit(); self.db.refresh(obj); return obj
    def update(self, obj: T, data: dict[str, Any]):
        for k, v in data.items(): setattr(obj, k, v)
        self.db.commit(); self.db.refresh(obj); return obj
    def delete(self, obj: T): self.db.delete(obj); self.db.commit()
