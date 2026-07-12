# Phase 5: Data & Persistence

**Estimated Duration:** 3-4 weeks  
**Prerequisites:** Phase 4 (I/O, Networking & Web Development)  
**Focus:** JSON parsing, raw SQL database operations via standard library, GORM ORM integration, and other binary serialization formats (MessagePack, Protobuf, XML).

---

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- [ ] Marshal and unmarshal JSON structures, customize tag mappings, handle dynamic payloads with maps, and write custom date serialization handlers.
- [ ] Connect to databases, leverage standard pool limits (`SetMaxOpenConns`, `SetMaxIdleConns`), prevent SQL Injection, run prepared statements, and control transaction rollbacks under errors.
- [ ] Incorporate GORM ORM, design entity models, handle One-to-One, One-to-Many, and Many-to-Many relationships, utilize AutoMigrate, and execute atomic transactions.
- [ ] Marshal/unmarshal XML data structures with attribute mapping.
- [ ] Handle binary serialization formats like MessagePack and understand Protocol Buffers' varint encoding and wire-level operations.

---

## 📚 Topics Covered

### 1. JSON Processing (`01-json-processing`)
- **[basic_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/basic_json.go)**: Standard marshaling/unmarshaling, pretty printing, and struct tagging.
- **[custom_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/custom_json.go)**: Implementing `json.Marshaler` and `json.Unmarshaler` for custom date formats.
- **[nested_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/nested_json.go)**: Multi-level objects, arrays, and map unmarshaling.
- **[performance.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/performance.go)**: Encoders/Decoders for streams, and UseNumber precision tuning.
- **[practice.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/practice.go)**: Snake/camel mapping configurations and currency string parsers.

### 2. Database Fundamentals (`02-database-fundamentals`)
- **[connection.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/connection.go)**: Standard SQLite connections, connection pool configs, and pings.
- **[queries.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/queries.go)**: Parameterized SQL inserts/selects, looping over rows safely, and prepared statements.
- **[transactions.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/transactions.go)**: Atomic transaction boundaries with deferred safety rollbacks.
- **[practice.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/practice.go)**: E-commerce accounts ledger and multi-account balance transactions.

### 3. ORMs (`03-orm`)
- **[models.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/models.go)**: GORM model schemas, foreign keys, constraints, and many-to-many associations.
- **[connection_orm.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/connection_orm.go)**: GORM connection options, auto-migration, and log formatting.
- **[crud.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/crud.go)**: Eager loading via Preload, partial updates, and soft/hard deletes.
- **[transactions_orm.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/transactions_orm.go)**: Auto-committing GORM transaction scopes.
- **[practice.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/practice.go)**: Customer-Order-OrderItem schema migrations and checkout calculations.

### 4. Data Serialization (`04-serialization`)
- **[xml.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/xml.go)**: XML elements, nested XML paths, and custom attributes.
- **[msgpack.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/msgpack.go)**: MsgPack binary serialization for caching.
- **[protobuf.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/protobuf.go)**: Protobuf varint bitwise encoding, field keys, and wire format unmarshaling.
- **[practice.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/practice.go)**: Config XML export and MsgPack cache cachers.

---

## 🚀 How to Run the Code

```bash
# Go to the Data & Persistence folder
cd phase-5-data-persistence/01-json-processing

# Run the tests to verify practice exercises
go test -v ./...
```
