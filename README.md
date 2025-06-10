# 쿠폰 발행 시스템

## 1. 개요
* [과제 내용](./doc/PROBLEM.md)

## 2. 기술 선택
- **connectrpc**: 과제 기본 요구사항

### 2.1 고려 사항

## 3. 아키텍처
클린 아키텍처 + DDD를 적용하고 있지만 의미없는 계층 구분에 추가적인 비용을 사용하지 않고 어디에서나 통용될 수 있는 아키텍처를 구성하는 것에 집중했습니다.
- 서로 분리된 레이어는 리포지토리를 통해 정보를 교환한다.
- 데이터 영속성 계층을 분리하는 경우 외에는 불필요한 인터페이스 추상화를 지양한다.
- 비즈니스 요구사항은 하나의 독립적인 기능이어야 한다. 비즈니스 로직의 변경이 의도치 않은 다른 유스케이스의 변동으로 전파되지 않도록 주의한다.
- DTO를 작성하는 것도 비용이기 때문에 유스케이스 구조체 내부에 Input/Output 구조체를 정의하고, Execute 메소드는 단일 Input을 매개변수로 받고 단일 Output을 반환한다.

### 3.1. 파일트리
기본적으로 클린 아키텍처를 기반으로 작성했으며, DDD에서 권장하는 컨벤션에 따라 아래와 같은 레이어를 가집니다.

- **domain**: 비즈니스 규칙과 데이터 모델과 연관된 엔티티, 값 객체, 열거형, 리포지토리, 도메인 서비스 등이 여기에 위치합니다.
- **application**: 비즈니스 로직을 담당하는 유스케이스가 여기에 위치합니다.
- **infra**: 서드파티 등 서비스와 느슨한 연결점을 가진 인프라스트럭처가 여기에 위치합니다.
- **interfaces**: 외부와 상호작용하는 엔드포인트가 여기에 위치합니다.

### 3.2. 시퀀스 다이어그램
```mermaid
sequenceDiagram
    title 새로운 쿠폰 캠페인 생성 (CreateCampaign)
    autonumber

    actor A as Admin
    participant B as Handler
    participant C as UseCase
    participant D as Repository
    participant E as Persistence

    A->>+B: 쿠폰 캠페인 생성 요청
    B->>+C: CreateCampaignUC 실행
    deactivate B
    C->>+D: 주어진 정보를 바탕으로 캠페인 생성 요청
    deactivate C
    D->>E: 영속성에 캠페인 저장 요청
    deactivate D
    E-->>+D: 캠페인 저장 완료
    D-->>+C: 캠페인 생성 완료 반환
    deactivate D
    C-->>+B: (campaign_create.Output, nil) 반환
    deactivate C
    B-->>A: 쿠폰 캠페인 생성 완료 (status: 201)
    deactivate B
```

```mermaid
sequenceDiagram
    title 쿠폰코드를 포함한 캠페인 정보 조회 (GetCampaign)
    autonumber

    actor A as User
    participant B as Handler
    participant C as UseCase
    participant D as Repository
    participant E as Persistence

    A->>+B: 쿠폰 및 캠페인 정보 요청
    B->>+C: GetCampaignUC 실행
    deactivate B
    C->>+D: 캠페인 ID에 매칭되는 정보 조회
    deactivate C
    D->>E: 영속성에서 캠페인 ID로 엔티티 조회
    deactivate D
    alt 엔티티 존재
        E-->>+D: 쿠폰 및 캠페인 정보 전달
        D-->>+C: 쿠폰 및 캠페인 정보 반환
        deactivate D
        C-->>+B: (campaign_get.Output, nil) 반환
        deactivate C
        B-->>A: 쿠폰 및 캠페인 정보 조회 완료 (status: 200)
        deactivate B
    else 엔티티 없음
        E-->>+D: 캠페인 정보 없음 전달
        D-->>+C: 캠페인 정보 없음 반환
        deactivate D
        C-->>+B: (nil, error) 반환
        deactivate C
        B-->>A: 쿠폰 및 캠페인 정보 조회 실패 (status: 404)
        deactivate B
    end
```

```mermaid
sequenceDiagram
    title 캠페인에서 쿠폰 발행 (IssueCoupon)
    autonumber

    actor A as User
    participant B as Handler
    participant C as UseCase
    participant D as Repository
    participant E as Persistence

    A->>+B: 선택한 캠페인에서 쿠폰 발행 요청
    B->>+C: IssueCouponUC 실행
    deactivate B
    C->>C: 트랜잭션 열기 (Read Committed)
    C->>+D: 캠페인 ID에 매칭되는 수정을 위한 정보 조회
    deactivate C
    D->>E: 영속성에서 캠페인 ID로 엔티티 조회 (비관적 락)
    deactivate D
    alt 엔티티 없음
        E-->>+D: 쿠폰 및 캠페인 정보 없음 전달
        D-->>+C: 쿠폰 및 캠페인 정보 없음 반환
        deactivate D
        C-->>C: 트랜잭션 롤백
        C-->>+B: (nil, error) 반환
        deactivate C
        B-->>A: "캠페인이 존재하지 않음" 반환 (status: 404)
        deactivate B
    else 엔티티 존재
        E-->>+D: 쿠폰 및 캠페인 정보 전달
        D-->>+C: 쿠폰 및 캠페인 정보 반환
        deactivate D
        C-->>C: 캠페인 시작일시 확인
        alt 캠페인 시작일시 이전
            C-->>C: 트랜잭션 롤백
            C-->>+B: (nil, error) 반환
            B-->>A: "캠페인이 아직 시작되지 않음" 반환 (status: 422)
            deactivate B
        end
        C-->>C: 잔여 쿠폰 수 확인
        alt 발행 가능한 쿠폰 없음
            C-->>C: 트랜잭션 롤백
            C-->>+B: (nil, error) 반환
            B-->>A: "잔여 쿠폰 없음" 반환 (status: 409)
            deactivate B
        else 잔여 쿠폰 수 > 0
            C-->>C: 쿠폰 수량 1개 차감
            C->>+D: 수정된 캠페인 정보 저장 요청
            deactivate C
            D->>E: 영속성에 캠페인 저장 요청
            deactivate D
            E-->>+D: 캠페인 저장 완료
            D-->>+C: 캠페인 저장 완료 반환
            deactivate D
            C-->>C: 10자리 한글과 숫자로 구성된 난수 생성
            C->>+D: 난수를 포함한 쿠폰 정보 저장 요청
            deactivate C
            D->>E: 영속성에 쿠폰 저장 요청
            deactivate D
            E-->>+D: 쿠폰 저장 완료
            D-->>+C: 쿠폰 저장 완료 반환
            deactivate D
            C-->>C: 트랜잭션 커밋
            C-->>+B: (coupon_issue.Output, nil) 반환
            deactivate C
            B-->>A: "쿠폰 발행 성공" 반환 (status: 201)
            deactivate B
        end
    end
```

### 3.3. ERD

## 4. 핵심 문제 해결 전략

### 4.1. 문제 해석
시작 부터 의문점이 생겼다. "쿠폰 발행은 지정된 일시에 자동으로 시작되어야 한다"는 문구는 단순 엔드포인트를 작성하는 것을 넘어 마치 프로덕션에 E2E 테스트케이스가 포함되어 있는 구조를 떠올리게 했다. 출고되자마자 갑자기 365일 24시간 공회전을 하기 시작하는 자동차를 상상해 봤다. 단지 백그라운드 작업을 구성할 수 있는 능력을 시험하는 것이 출제 의도였을까?

우선 캠페인 생성이라는 엔드포인트가 요구사항이라는 점에 주목해 보기로 했다. 캠페인 생성 시에 시작일시를 정하고 생성할 텐데, 이 시작일시를 기준으로 스케줄러를 시작하면 될 것 같다. 따라서 cmd 레이어에는 컴파일을 위한 main.go 파일이 두 개가 되어야 한다.

### 4.2. 동시성 제어

### 4.3. 기타 해결 전략

## 5. 실행 및 테스트 방법

### 5.1. 시스템 실행

### 5.2. 부하 테스트 실행

## 6. 테스트 결과 및 결론

## 7. 향후 개선 과제
